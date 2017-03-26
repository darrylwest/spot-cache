//
// Request - message request structure and functions
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-19 10:16:38
//

package spotcache

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/oklog/ulid"
	"io"
	"math/rand"
	"time"
)

var (
	entropy  io.Reader = rand.New(rand.NewSource(time.Now().UnixNano()))
	zerobyte           = make([]byte, 0)
)

func genulid(entropy io.Reader, ts uint64) (ulid.ULID, error) {
	value, err := ulid.New(ts, entropy)
	return value, err
}

func CreateRawULID() ulid.ULID {
	ts := uint64(time.Now().UnixNano() / 1000000)
	v, _ := genulid(entropy, ts)

	return v
}

// generate and return a ulid as a string
func CreateULID() string {
	return CreateRawULID().String()
}

type IdType [26]byte
type SessionType [12]byte

// request object as created by the client
type Request struct {
	Id       IdType
	Session  SessionType
	Op       CommandOp
	MetaSize uint16
	KeySize  uint16
	DataSize uint32
	Metadata []byte
	Key      []byte
	Value    []byte
}

type RequestBuilder struct {
	session SessionType
}

func NewRequestBuilder(sess SessionType) *RequestBuilder {
	b := RequestBuilder{session: sess}

	return &b
}

func CreateCommandId() IdType {
	var id IdType
	copy(id[:26], []byte(CreateULID()))

	return id
}

func (rb RequestBuilder) NewRequest(op CommandOp) Request {
	req := Request{}

	// create the request id...
	req.Id = CreateCommandId()
	req.Session = rb.session
	req.Op = op

	return req
}

func (req *Request) updateRequest(key, value, metadata []byte) {
	req.MetaSize = uint16(len(metadata))
	req.KeySize = uint16(len(key))
	req.DataSize = uint32(len(value))

	req.Metadata = metadata
	req.Key = key
	req.Value = value
}

// create a put command with the current session
func (rb *RequestBuilder) CreatePutRequest(key, value, metadata []byte) *Request {
	req := rb.NewRequest(PUT)

	req.updateRequest(key, value, metadata)

	return &req
}

func (rb *RequestBuilder) CreateGetRequest(key, metadata []byte) *Request {
	req := rb.NewRequest(GET)

	req.updateRequest(key, zerobyte, metadata)

	return &req
}

func (rb *RequestBuilder) CreateHasRequest(key, metadata []byte) *Request {
	req := rb.NewRequest(HAS)

	req.updateRequest(key, zerobyte, zerobyte)

	return &req
}

func (rb *RequestBuilder) CreateDeleteRequest(key, metadata []byte) *Request {
	req := rb.NewRequest(DELETE)

	req.updateRequest(key, zerobyte, zerobyte)

	return &req
}

// decode the little endian bytes and parse into request object
func RequestFromBytes(buf []byte) (*Request, error) {
	raw := bytes.NewReader(buf)
	ba := make([]byte, len(buf))

	req := Request{}
	// this may be unnecessary if the socket reader does the decoding...
	err := binary.Read(raw, binary.LittleEndian, ba)

	sz := len(req.Id)
	idx, idy := 0, sz

	copy(req.Id[0:sz], ba[idx:idy])

	sz = len(req.Session)
	idx, idy = idy, idy+sz
	copy(req.Session[0:sz], ba[idx:idy])

	sz = 1
	idx, idy = idy, idy+sz
	req.Op = CommandOp(ba[idx])

	sz = 2
	idx, idy = idy, idy+sz
	req.MetaSize = uint16(ba[idx:idy][0])

	idx, idy = idy, idy+sz
	req.KeySize = uint16(ba[idx:idy][0])

	sz = 4
	idx, idy = idy, idy+sz
	req.DataSize = uint32(ba[idx:idy][0])

	idx, idy = idy, idy+int(req.MetaSize)
	req.Metadata = ba[idx:idy]

	idx, idy = idy, idy+int(req.KeySize)
	req.Key = ba[idx:idy]

	idx, idy = idy, idy+int(req.DataSize)
	req.Value = ba[idx:idy]

	return &req, err
}

// encode the request into a stream of little endian bytes; return error if encoding fails
func (req *Request) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		req.Id,
		req.Session,
		req.Op,
		req.MetaSize,
		req.KeySize,
		req.DataSize,
		req.Metadata,
		req.Key,
		req.Value,
	}

	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			log.Error("encoding error %v", err)
			return []byte(""), err
		}
	}

	return buf.Bytes(), nil
}

func (req *Request) String() string {
	return fmt.Sprintf(
		"Id:%s,Session:%s,Op:%d,MetaSize:%d,KeySize:%d,DataSize:%d,Metadata:%s,Key:%s,Value:%v",
		req.Id, req.Session, req.Op,
		req.MetaSize, req.KeySize, req.DataSize,
		req.Metadata, req.Key, req.Value)
}
