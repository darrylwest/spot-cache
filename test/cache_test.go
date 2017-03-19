//
// CacheTest - test the get/put/has/delete methods
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-18 13:53:58
//

package test

import (
	// "fmt"
	"reflect"
	"spotcache"
	"testing"

	. "github.com/franela/goblin"
)

func TestCache(t *testing.T) {
	g := Goblin(t)

	g.Describe("Cache", func() {
		cfg := spotcache.NewConfigForEnvironment("test")
		spotcache.CreateLogger(cfg)

		// before
		// after

		g.It("should create a cache object", func() {
			cache := spotcache.NewCache(cfg)

			g.Assert(reflect.TypeOf(cache).String()).Equal("*spotcache.Cache")
			// test stuff...
		})
	})
}
