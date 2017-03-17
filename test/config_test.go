package test

import (
	"os"
	"spotcache"
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

			g.Assert(cfg.GetUnixSock()).Equal("")
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := spotcache.NewDefaultConfig()

			g.Assert(cfg.GetUnixSock()).Equal(home + "/spot.sock")

			hash := cfg.ToMap()

			g.Assert(hash != nil)

			if value, ok := hash["webroot"]; ok {
				g.Assert(value).Equal("public")
			}

			g.Assert(hash["home"]).Equal(home)
			g.Assert(hash["baseport"]).Equal(3001)
			g.Assert(hash["unixsock"]).Equal(home + "/spot.sock")
			g.Assert(hash["timeout"]).Equal(int64(600))
		})

		g.It("should create context from args", func() {
			cfg := spotcache.ParseArgs()

			hash := cfg.ToMap()

			g.Assert(hash["home"]).Equal(home)
			g.Assert(hash["baseport"]).Equal(3001)
			g.Assert(hash["unixsock"]).Equal(home + "/spot.sock")
			g.Assert(hash["timeout"]).Equal(int64(600))
		})
	})
}
