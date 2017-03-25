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

var entropy io.Reader = rand.New(rand.NewSource(time.Now().UnixNano()))

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

// create a put command with the current session
func (rb *RequestBuilder) CreatePutCommand(key, value, metadata []byte) *Request {
	req := Request{}

	// create the request id...
	copy(req.Id[:26], []byte(CreateULID()))
	req.Session = rb.session
	req.Op = PUT
	req.MetaSize = uint16(len(metadata))
	req.KeySize = uint16(len(key))
	req.DataSize = uint32(len(value))

	req.Metadata = metadata
	req.Key = key
	req.Value = value

	return &req
}

/*
func RequestFromBytes(buf []byte) (*Request, error) {
    req := new(Request)
    br := bytes.NewReader(buf)
    fmt.Println(br)
    if err := binary.Read(br, binary.LittleEndian, req.Id); err != nil {
        return nil, err
    }

    if err := binary.Read(br, binary.LittleEndian, req.Session); err != nil {
        return nil, err
    }

    return req, nil
}
*/

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
