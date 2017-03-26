//
// @RequestTest
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-19 12:09:24
//

package test

import (
	"reflect"
	"testing"

	"spotcache"

	. "github.com/franela/goblin"
)

func TestRequest(t *testing.T) {
	g := Goblin(t)

	g.Describe("Request", func() {
		cfg := spotcache.NewConfigForEnvironment("test")
		spotcache.CreateLogger(cfg)
		var session spotcache.SessionType
		copy(session[:12], []byte(spotcache.CreateSessionId()))

		builder := spotcache.NewRequestBuilder(session)
		key := []byte(spotcache.CreateULID())
		value := CreateRandomData()
		metadata := []byte("expire:40;")

        keysz := uint16(len(key))
        metasz := uint16(len(metadata))
        valsz := uint32(len(value))

        g.It("should create a new instance of Request from NewCommand", func() {
            req := builder.NewRequest(spotcache.NOOP)

			g.Assert(reflect.TypeOf(req).String()).Equal("spotcache.Request")
			g.Assert(len(req.Id)).Equal(26)
			g.Assert(len(req.Session)).Equal(12)
            g.Assert(req.Op).Equal(spotcache.NOOP)
        })

		g.It("should create a put instance of Request", func() {
			req := builder.CreatePutRequest(key, value, metadata)

			g.Assert(reflect.TypeOf(req).String()).Equal("*spotcache.Request")
			g.Assert(req.Op).Equal(spotcache.PUT)
            g.Assert(req.MetaSize).Equal(metasz)
            g.Assert(req.KeySize).Equal(keysz)
            g.Assert(req.DataSize).Equal(valsz)

            g.Assert(req.Metadata).Equal(metadata)
            g.Assert(req.Key).Equal(key)
            g.Assert(req.Value).Equal(value)
		})

		g.It("should create a get instance of Request", func() {
            req := builder.CreateGetRequest(key, metadata)

			g.Assert(reflect.TypeOf(req).String()).Equal("*spotcache.Request")
			g.Assert(req.Op).Equal(spotcache.GET)
            g.Assert(req.MetaSize).Equal(metasz)
            g.Assert(req.KeySize).Equal(keysz)
            g.Assert(req.DataSize).Equal(uint32(0))

            g.Assert(req.Metadata).Equal(metadata)
            g.Assert(req.Key).Equal(key)
            g.Assert(req.Value).Equal(make([]byte, 0))
        })

		g.It("should create a has request")
		g.It("should create a delete request")
		g.It("should create an expire request")
		g.It("should create a ttl request")
		g.It("should create a subscribe request")
		g.It("should create an unsubscribe request")
		g.It("should create a publish request")
		g.It("should create a status request")
		g.It("should create a ping request")
		g.It("should create a shutdown request")

		g.It("should create an instance of RequestBuilder")
		g.It("should create a byte stream from a request object")
		g.It("should read and parse a byte stream into a request object")
	})
}
