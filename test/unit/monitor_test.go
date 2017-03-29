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
	"os"
	"time"

	. "github.com/franela/goblin"
)

func TestMonitor(t *testing.T) {
	g := Goblin(t)

	g.Describe("Monitor", func() {
		// before?
		// after?
		cfg := spotcache.NewConfigForEnvironment("test")
		spotcache.CreateLogger(cfg)

		g.It("should create a new monitor struct", func() {
			monitor := spotcache.NewMonitorService(cfg)

			g.Assert(monitor.Sockfile).Equal(cfg.Unixsock)
			g.Assert(monitor.CreateDate.IsZero()).IsTrue("should be a zero date")
		})

		g.It("should open and serve then close a unix socket when messaged to stop", func(done Done) {
			monitor := spotcache.NewMonitorService(cfg)

			stop := make(chan bool)
			go func() {
				time.Sleep(time.Millisecond * 10)
				info, err := os.Stat(monitor.Sockfile)

				// fmt.Println(info)

				g.Assert(err == nil).IsTrue("error should be nil")
				g.Assert(info.Name()).Equal("spot.sock")

				stop <- true
			}()

			monitor.OpenAndServe(stop)

			// just wait for it to get the stop message and insure that the createdate is not zero...
			g.Assert(monitor.CreateDate.Year()).Equal(time.Now().UTC().Year())
			done()
		})

		g.It("should parse a monitor command")
		g.It("should handle a shutdown command")
	})
}
