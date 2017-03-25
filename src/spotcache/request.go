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

// request object as created by the client
type Request struct {
	Id       []byte
	Session  []byte
	Op       CommandOp
	MetaSize uint16
	KeySize  uint16
	DataSize uint32
	Metadata []byte
	Key      []byte
	Value    []byte
}

type RequestBuilder struct {
	session []byte
}

func NewRequestBuilder(sess []byte) *RequestBuilder {
	b := RequestBuilder{session: sess}

	return &b
}

// create a put command with the current session
func (rb *RequestBuilder) CreatePutCommand(key, value, metadata []byte) *Request {
	req := Request{}

	// create the request id...
	req.Id = []byte(CreateULID())
	req.Session = rb.session
	req.Op = []byte("pu")
	req.MetaSize = uint16(len(metadata))
	req.KeySize = uint16(len(key))
	req.DataSize = uint32(len(value))

	req.Metadata = metadata
	req.Key = key
	req.Value = value

	return &req
}

// encode the request into a stream of little endian bytes; return error if encoding fails
func (req *Request) ToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	var data = []interface{}{}

	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			log.Error("encoding error %v", err)
			return []byte(""), err
		}
	}

	return buf.Bytes(), nil
}
