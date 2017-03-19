// test fixtures; should move this to it's own repo for random-utils or something
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-12 11:39:18

package test

import (
	"fmt"
	"github.com/oklog/ulid"
	"io"
	"math/rand"
	"time"

	"spotcache"
)

type RequestBuilder struct {
	session []byte
}

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

func CreateRandomId() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%x%x", r.Intn(9e7)+1e8, r.Intn(9e7)+1e8)
}

// create a command id (ulid)
func CreateCommandId() []byte {
	return []byte(CreateULID())
}

func CreateRandomData() []byte {
	return []byte(CreateRandomId() + "..." + CreateULID())
}

// TODO : create helper "CommandFactory" "RequestBuilder", etc methods that will be
// eventually ported to the client application

func NewRequestBuilder(sess []byte) *RequestBuilder {
	b := RequestBuilder{session: sess}

	return &b
}

// create a valid request object with a put op, random key and value
func (b *RequestBuilder) CreatePutCommand(key, value, metadata []byte) spotcache.Request {
	req := spotcache.Request{}

	req.Id = CreateCommandId()
	req.Session = b.session
	req.Op = []byte("pu")
	req.MetaSize = uint16(len(metadata))
	req.KeySize = uint16(len(key))
	req.DataSize = uint32(len(value))
	req.Metadata = metadata
	req.Key = key
	req.Value = value

	return req
}
