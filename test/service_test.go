package test

import (
	"spotcache"
	"testing"
	// "fmt"
	"time"

	. "github.com/franela/goblin"
)

func TestService(t *testing.T) {
	g := Goblin(t)

	g.Describe("Service", func() {
		// before?
		// after?
		cfg := spotcache.NewConfigForEnvironment("test")
		spotcache.CreateLogger(cfg)

		g.It("should create a new cache service object", func() {
			service := spotcache.NewCacheService(cfg)

			g.Assert(service.Port).Equal(cfg.Baseport)
			g.Assert(service.CreateDate.IsZero()).IsTrue("should be a zero date")
			g.Assert(service.ClientCount).Equal(0)
		})

		g.It("should open and serve then close the service", func(done Done) {
			service := spotcache.NewCacheService(cfg)

			stop := make(chan bool)

			go func() {
				time.Sleep(time.Millisecond * 10)

				stop <- true
			}()

			service.OpenAndServe(stop)

			g.Assert(service.CreateDate.Year()).Equal(time.Now().UTC().Year())

			done()
		})

		g.It("should handle a client connection shutdown on error")
		g.It("should start a client session with session id")
	})
}
