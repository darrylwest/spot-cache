package unit

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
			g.Assert(cfg.Baseport).Equal(0)
			g.Assert(cfg.Home).Equal("")
		})

		g.It("should create a context struct with defaults set", func() {
			cfg := spotcache.NewDefaultConfig()

			g.Assert(cfg.Home).Equal(home)
			g.Assert(cfg.Baseport).Equal(19501)
			g.Assert(cfg.Timeout).Equal(int64(600))
			g.Assert(len(cfg.Logpath) > 0).IsTrue()
			g.Assert(len(cfg.Logname) > 0).IsTrue()
			g.Assert(len(cfg.Dbpath) > 0).IsTrue()
		})
	})
}
