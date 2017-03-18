// simple verion
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 09:59:37

package spotcache

import "fmt"

const (
	major = 0
	minor = 90
	patch = 105
)

func Version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
