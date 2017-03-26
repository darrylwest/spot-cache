//
// CacheTest - test the get/put/has/delete methods; other cache tests exist in the command_test.go file
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
		cache := spotcache.NewCache(cfg)

		g.Before(func() {
			cache.Open()
		})

		g.After(func() {
			cache.Close()
		})

		g.It("should create a cache object", func() {
			g.Assert(reflect.TypeOf(cache).String()).Equal("*spotcache.Cache")
			// test stuff...
		})

		g.It("should put data with expire/ttl")
		g.It("should get data that has not expired and update ttl")
		g.It("should attempt get with null return after data expires")
		g.It("should return true if data exists and update ttl")
		g.It("should return false if data does not exist")
		g.It("should return all keys for the current cache", func() {
			keys, _ := cache.Keys()

			g.Assert(len(keys) > 0).IsTrue()
		})

        // Put

        // Get

        // Has
        g.It("should return true from Has if a key exists")
        g.It("should return false from Has if a key does not exist")

        // Delete
	})
}
