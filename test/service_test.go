package test

import (
	"spotcache"
	"testing"

	. "github.com/franela/goblin"
)

func TestService(t *testing.T) {
	g := Goblin(t)

	g.Describe("Service", func() {
		spotcache.CreateLogger(spotcache.NewConfigForEnvironment("test"))

		g.It("should handle a client shutdown request")
		g.It("should handle a client connection shutdown on error")
        g.It("should start a client session with session id")
	})
}
