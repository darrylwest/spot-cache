// test fixtures; should move this to it's own repo for random-utils or something
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-12 11:39:18

package unit

import (
	"fmt"
	"github.com/darrylwest/go-unique/unique"
)

// CreateULID generate and return a ulid as a string
func CreateULID() string {
	return unique.CreateULID()
}

// CreateRandomID create a random id
func CreateRandomID() string {
	b1, _ := unique.RandomBytes(12)
	b2, _ := unique.RandomBytes(16)
	return fmt.Sprintf("%x%x", b1, b2)
}

// CreateRandomData create random data
func CreateRandomData() []byte {
	return []byte(CreateRandomID() + "..." + CreateULID())
}
