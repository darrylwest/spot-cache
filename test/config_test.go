package webservertest

import (
	"testing"
	"service"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		webserver.CreateLogger(webserver.NewConfigForEnvironment("test"))

		g.It("should create a context struct", func() {
			ctx := new(webserver.Config)

			g.Assert(ctx.GetShutdownPort()).Equal(0)
		})

		g.It("should create a context struct with defaults set", func() {
			ctx := webserver.NewDefaultConfig()

			g.Assert(ctx.GetShutdownPort()).Equal(3009)

			hash := ctx.ToMap()

			g.Assert(hash != nil)

			if value, ok := hash["webroot"]; ok {
				g.Assert(value).Equal("public")
			}

			g.Assert(hash["baseport"]).Equal(3001)
			g.Assert(hash["shutdownPort"]).Equal(3009)
			g.Assert(hash["serverCount"]).Equal(2)
			g.Assert(hash["timeout"]).Equal(int64(600))
		})

		g.It("should create context from args", func() {
			ctx := webserver.ParseArgs()

			g.Assert(ctx.GetShutdownPort()).Equal(3009)

			hash := ctx.ToMap()

			g.Assert(hash["webroot"]).Equal("public")
			g.Assert(hash["baseport"]).Equal(3001)
			g.Assert(hash["shutdownPort"]).Equal(3009)
			g.Assert(hash["serverCount"]).Equal(2)
			g.Assert(hash["timeout"]).Equal(int64(600))
		})

	})
}
