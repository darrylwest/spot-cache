package spotcache_test

import (
	"testing"
	"spotcache"

	. "github.com/franela/goblin"
)

func TestService(t *testing.T) {
	g := Goblin(t)

	g.Describe("Service", func() {
        spotcache.CreateLogger(spotcache.NewConfigForEnvironment("test"))

		g.It("should start a mock service")
		g.It("should handle a mock client connection")
	})
}
