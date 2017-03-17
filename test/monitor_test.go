//
// MonitorTest
//
// @author darryl west <darryl.west@raincitysoftware.com>
// @created 2017-03-17 08:32:48
//

package test

import (
	// "spotcache"
	"testing"
	// "fmt"

	. "github.com/franela/goblin"
)

func TestMonitor(t *testing.T) {
	g := Goblin(t)

	g.Describe("Monitor", func() {
		// before?
		// after?

		g.It("should open and close a unix socket")
		g.It("should parse a monitor command")
		g.It("should handle a shutdown command")
	})
}
