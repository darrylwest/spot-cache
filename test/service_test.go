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
			g.Assert(service.Timeout).Equal(time.Duration(1e9))
		})

        g.It("should create a server socket listener", func() {
			service := spotcache.NewCacheService(cfg)
            service.Port = 4000
			ss, err := service.CreateListener()
			g.Assert(err).Equal(nil)
            g.Assert(ss.Addr().String()).Equal("127.0.0.1:4000")
        })

		g.It("should open and serve then close the service", func(done Done) {
			service := spotcache.NewCacheService(cfg)
			service.Timeout = time.Duration(1e6)

			ss, err := service.CreateListener()
			g.Assert(err).Equal(nil)

			go func() {
				service.ListenAndServe(ss)
				g.Assert(service.CreateDate.Year()).Equal(time.Now().UTC().Year())
			}()

			time.Sleep(time.Millisecond * 100)
			service.Shutdown()

			done()
		})

		g.It("should handle a client connection shutdown on error")
		g.It("should start a client session with session id")
	})
}
