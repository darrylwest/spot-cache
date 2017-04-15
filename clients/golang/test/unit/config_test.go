//
// ConfigTest - test the client configuration
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-15 08:59:14
//

package unit

import (
    "os"
    "spotclient"
    "testing"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
    g := Goblin(t)

    g.Describe("Config", func() {
		spotcache.CreateLogger(spotcache.NewConfigForEnvironment("test"))
		home := os.Getenv("HOME") + "/.spotcache"

		g.It("should create a config struct", func() {
			cfg := new(spotcache.Config)

			g.Assert(cfg.Unixsock).Equal("")
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := spotcache.NewDefaultConfig()

			g.Assert(cfg.Unixsock).Equal(home + "/spot.sock")

			g.Assert(cfg.Home).Equal(os.Getenv("HOME") + "/.spotcache")
			g.Assert(cfg.Baseport).Equal(3001)
			g.Assert(cfg.Unixsock).Equal(cfg.Home + "/spot.sock")
			g.Assert(cfg.Timeout).Equal(int64(600))
		})
    })
}

