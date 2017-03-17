//
// MonitorTest
//
// @author darryl west <darryl.west@raincitysoftware.com>
// @created 2017-03-17 08:32:48
//

package test

import (
	"spotcache"
	"testing"
	// "fmt"
    "time"

	. "github.com/franela/goblin"
)

func TestMonitor(t *testing.T) {
	g := Goblin(t)

	g.Describe("Monitor", func() {
		// before?
		// after?
        cfg := spotcache.NewConfigForEnvironment("test")
        cmap := cfg.ToMap()

        g.It("should create a new monitor struct", func() {
            monitor := spotcache.NewMonitor(cfg)
            
            g.Assert(monitor.Sockfile).Equal(cmap["unixsock"])
            g.Assert(monitor.CreateDate.Year()).Equal( time.Now().UTC().Year() )
        });

		g.It("should open and close a unix socket")
		g.It("should parse a monitor command")
		g.It("should handle a shutdown command")
	})
}

