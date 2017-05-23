//
// Config - the client config
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-14 09:36:57
//

package clienttest

import (
	"spotclient"
	"testing"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("ClientConfig", func() {

		g.It("should create a config struct", func() {
			cfg := spotclient.NewConfigForEnvironment("test")

			g.Assert(cfg != nil).IsTrue()
			g.Assert(cfg.Env).Equal("test")
			g.Assert(cfg.Host).Equal("localhost")
			g.Assert(cfg.Port).Equal(19501)
			g.Assert(cfg.Timeout).Equal(int64(600))
		})

		g.It("should parse the command line args and return a config object", func() {
			cfg := spotclient.ParseArgs()
			g.Assert(cfg != nil).IsTrue()

			g.Assert(cfg != nil).IsTrue()
			g.Assert(cfg.Env).Equal("production")
			g.Assert(cfg.Host).Equal("localhost")
			g.Assert(cfg.Port).Equal(19501)
			g.Assert(cfg.Timeout).Equal(int64(600))
		})
	})
}
