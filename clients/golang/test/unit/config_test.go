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
		home := os.Getenv("HOME") + "/.spotcache"

		g.It("should create a config struct", func() {
			cfg := new(spotclient.Config)

			g.Assert(cfg.Port).Equal(0)
            g.Assert(cfg.Home).Equal("")
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := spotclient.NewDefaultConfig()

            g.Assert(cfg.Home).Equal(home)
			g.Assert(cfg.Port).Equal(3001)
			g.Assert(cfg.Timeout).Equal(int64(600))
		})
    })
}

