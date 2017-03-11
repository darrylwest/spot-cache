package spotcache_test

import (
	"testing"
	"spotcache"

	. "github.com/franela/goblin"
)

func TestCommand(t *testing.T) {
	g := Goblin(t)

	g.Describe("Command", func() {
        spotcache.CreateLogger(spotcache.NewConfigForEnvironment("test"))

		g.It("should parse a put command")
		g.It("should execute a put command")

		g.It("should parse a get command")
		g.It("should execute a get command")

		g.It("should parse a has command")
		g.It("should execute a has command")

		g.It("should parse a del command")
		g.It("should execute a del command")

		g.It("should parse a stat command")
		g.It("should execute a stat command")

		g.It("should parse a halt command")
		g.It("should execute a halt command")

	})
}
