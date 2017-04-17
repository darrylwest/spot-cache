//
// Request - message request structure and functions; this object isolates the command builder functions to make available to
//           both server and client as well as tests; this module also creates the Reponse object and associated utility functions
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

// CreateRawULID create a raw ulid
func CreateRawULID() ulid.ULID {
	ts := uint64(time.Now().UnixNano() / 1000000)
	v, _ := genulid(entropy, ts)

	return v
}

// CreateULID generate and return a ulid as a string
func CreateULID() string {
	return CreateRawULID().String()
}

// IDType 26 bytes
type IDType [26]byte

// SessionType - 12 bytes
type SessionType [12]byte

// Request request object as created by the client
type Request struct {
	ID       IDType
	Session  SessionType
	Op       CommandOp
	MetaSize uint16
	KeySize  uint16
	DataSize uint32
	Metadata []byte
	Key      []byte
	Value    []byte
}

// Response request object as created by the client
type Response struct {
	ID       IDType
	Session  SessionType
	Op       CommandOp
	MetaSize uint16
	DataSize uint32
	Metadata []byte
	Data     []byte
}

// RequestBuilder holds the session
type RequestBuilder struct {
	session SessionType
}

// NewRequestBuilder return a new request builder object
func NewRequestBuilder(sess SessionType) *RequestBuilder {
	b := RequestBuilder{session: sess}

	return &b
}

// CreateCommandID create a command id
func CreateCommandID() IDType {
	var id IDType
	copy(id[:26], []byte(CreateULID()))

	return id
}

// NewRequest create a new request object
func (rb RequestBuilder) NewRequest(op CommandOp) Request {
	req := Request{}

	// create the request id...
	req.ID = CreateCommandID()
	req.Session = rb.session
	req.Op = op

	return req
}

// updateRequest
func (req *Request) updateRequest(key, value, metadata []byte) {
	req.MetaSize = uint16(len(metadata))
	req.KeySize = uint16(len(key))
	req.DataSize = uint32(len(value))

	req.Metadata = metadata
	req.Key = key
	req.Value = value
}

// CreateResponse create a response object from the reqest, response value and new meta data
func (req *Request) CreateResponse(value, metadata []byte) *Response {
	resp := Response{}

	resp.ID = req.ID
	resp.Session = req.Session
	resp.Op = req.Op

	resp.MetaSize = uint16(len(metadata))
	resp.DataSize = uint32(len(value))

	resp.Metadata = metadata
	resp.Data = value

	return &resp
}

// CreatePutRequest create a put command with the current session
func (rb *RequestBuilder) CreatePutRequest(key, value, metadata []byte) *Request {
	req := rb.NewRequest(PUT)

	req.updateRequest(key, value, metadata)

	return &req
}

// CreateGetRequest create a get request
func (rb *RequestBuilder) CreateGetRequest(key, metadata []byte) *Request {
	req := rb.NewRequest(GET)

	req.updateRequest(key, zerobyte, metadata)

	return &req
}

// CreateHasRequest return the has request op
func (rb *RequestBuilder) CreateHasRequest(key, metadata []byte) *Request {
	req := rb.NewRequest(HAS)

	req.updateRequest(key, zerobyte, zerobyte)

	return &req
}

// CreateDeleteRequest create a delete request op
func (rb *RequestBuilder) CreateDeleteRequest(key, metadata []byte) *Request {
	req := rb.NewRequest(DELETE)

	req.updateRequest(key, zerobyte, zerobyte)

	return &req
}

// CreateKeysRequest create a keys request
func (rb *RequestBuilder) CreateKeysRequest(key, metadata []byte) *Request {
	req := rb.NewRequest(KEYS)

	req.updateRequest(key, zerobyte, zerobyte)

	return &req
}

// CreatePingRequest create a simple ping request
func (rb *RequestBuilder) CreatePingRequest() *Request {
	req := rb.NewRequest(PING)
	req.updateRequest(zerobyte, zerobyte, zerobyte)

	return &req
}

// RequestFromBytes decode the little endian bytes and parse into request object
func RequestFromBytes(buf []byte) (*Request, error) {
	raw := bytes.NewReader(buf)
	ba := make([]byte, len(buf))

	req := Request{}
	// this may be unnecessary if the socket reader does the decoding...
	err := binary.Read(raw, binary.LittleEndian, ba)

	sz := len(req.ID)
	idx, idy := 0, sz

	copy(req.ID[0:sz], ba[idx:idy])

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

// ResponseFromBytes decode the little endian bytes and parse into response object
func ResponseFromBytes(buf []byte) (*Response, error) {
	raw := bytes.NewReader(buf)
	ba := make([]byte, len(buf))

	res := Response{}
	// this may be unnecessary if the socket reader does the decoding...
	err := binary.Read(raw, binary.LittleEndian, ba)

	sz := len(res.ID)
	idx, idy := 0, sz

	copy(res.ID[0:sz], ba[idx:idy])

	sz = len(res.Session)
	idx, idy = idy, idy+sz
	copy(res.Session[0:sz], ba[idx:idy])

	sz = 1
	idx, idy = idy, idy+sz
	res.Op = CommandOp(ba[idx])

	sz = 2
	idx, idy = idy, idy+sz
	res.MetaSize = uint16(ba[idx:idy][0])

	sz = 4
	idx, idy = idy, idy+sz
	res.DataSize = uint32(ba[idx:idy][0])

	idx, idy = idy, idy+int(res.MetaSize)
	res.Metadata = ba[idx:idy]

	idx, idy = idy, idy+int(res.DataSize)
	res.Data = ba[idx:idy]

	return &res, err
}

// ToBytes encode the request into a stream of little endian bytes; return error if encoding fails
func (req *Request) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		req.ID,
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

// ToBytes encode the response into a stream of little endian bytes; return error if encoding fails
func (res *Response) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	var data = []interface{}{
		res.ID,
		res.Session,
		res.Op,
		res.MetaSize,
		res.DataSize,
		res.Metadata,
		res.Data,
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
		"ID:%s,Session:%s,Op:%d,MetaSize:%d,KeySize:%d,DataSize:%d,Metadata:%s,Key:%s,Value:%v",
		req.ID, req.Session, req.Op,
		req.MetaSize, req.KeySize, req.DataSize,
		req.Metadata, req.Key, req.Value)
}

func (res *Response) String() string {
	return fmt.Sprintf(
		"ID:%s,Session:%s,Op:%d,MetaSize:%d,DataSize:%d,Metadata:%s,Value:%v",
		res.ID, res.Session, res.Op,
		res.MetaSize, res.DataSize,
		res.Metadata, res.Data)
}
