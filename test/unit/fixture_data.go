// test fixtures; should move this to it's own repo for random-utils or something
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-12 11:39:18

package unit

import (
	"fmt"
	"github.com/oklog/ulid"
	"io"
	"math/rand"
	"time"
	// "spotcache"
)

var entropy io.Reader = rand.New(rand.NewSource(time.Now().UnixNano()))

// generate ulid with entropy and time
func genulid(entropy io.Reader, ts uint64) (ulid.ULID, error) {
	value, err := ulid.New(ts, entropy)
	return value, err
}

// CreateRawULID create the raw ULID
func CreateRawULID() ulid.ULID {
	ts := uint64(time.Now().UnixNano() / 1000000)
	v, _ := genulid(entropy, ts)

	return v
}

// CreateULID generate and return a ulid as a string
func CreateULID() string {
	return CreateRawULID().String()
}

// CreateRandomID create a random id
func CreateRandomID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%x%x", r.Intn(9e7)+1e8, r.Intn(9e7)+1e8)
}

// CreateRandomData create random data
func CreateRandomData() []byte {
	return []byte(CreateRandomID() + "..." + CreateULID())
}
