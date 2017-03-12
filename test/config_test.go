package test

import (
	"spotcache"
	"testing"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("Config", func() {
		spotcache.CreateLogger(spotcache.NewConfigForEnvironment("test"))

		g.It("should create a config struct", func() {
			cfg := new(spotcache.Config)

			g.Assert(cfg.GetShutdownPort()).Equal(0)
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := spotcache.NewDefaultConfig()

			g.Assert(cfg.GetShutdownPort()).Equal(3009)

			hash := cfg.ToMap()

			g.Assert(hash != nil)

			if value, ok := hash["webroot"]; ok {
				g.Assert(value).Equal("public")
			}

			g.Assert(hash["baseport"]).Equal(3001)
			g.Assert(hash["shutdownPort"]).Equal(3009)
			g.Assert(hash["timeout"]).Equal(int64(600))
		})

		g.It("should create context from args", func() {
			cfg := spotcache.ParseArgs()

			g.Assert(cfg.GetShutdownPort()).Equal(3009)

			hash := cfg.ToMap()

			g.Assert(hash["baseport"]).Equal(3001)
			g.Assert(hash["shutdownPort"]).Equal(3009)
			g.Assert(hash["timeout"]).Equal(int64(600))
		})
	})
}
