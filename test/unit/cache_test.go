//
// CacheTest - test the get/put/has/delete methods; other cache tests exist in the command_test.go file
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-18 13:53:58
//

package unit

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
		knownKey := []byte("myknownkey")
		knownValue := []byte("my known value")
		zeroTTL := spotcache.TTLSeconds(10)

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

		g.It("should put data with expire/ttl", func() {
			err := cache.Put(knownKey, knownValue, zeroTTL)
			g.Assert(err).Equal(nil)
		})

		g.It("should get data that has not expired and update ttl")
		g.It("should attempt get with null return after data expires")
		g.It("should return true if data exists and update ttl")
		g.It("should return false if data does not exist")
		g.It("should return all keys for the current cache", func() {
			keys, _ := cache.Keys()

			g.Assert(len(keys) > 0).IsTrue()
		})

		// Put
		g.It("should return ok after Put object", func() {
            err := cache.Put(knownKey, []byte("my new thing"), zeroTTL)
            g.Assert(err).Equal(nil)
        })

		// Get
		g.It("should return ok and an object from Get when the object exists", func() {
            val := []byte("this is my new value object ok?")
            cache.Put(knownKey, val, zeroTTL)
            value, err := cache.Get(knownKey)
            g.Assert(err).Equal(nil)
            g.Assert(value).Equal(val)
        })

		g.It("should return false and nil from Has if a object/key does not exist", func() {
            val, err := cache.Get([]byte("badbadkey"))
            g.Assert(err != nil).IsTrue()
            g.Assert(val).Equal([]byte(nil))
        })

		// Has
		g.It("should return true from Has if a key exists")
		g.It("should return false from Has if a key does not exist")

		// Delete
		g.It("should return ok if item was deleted")
	})
}
