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
		key := []byte("mytestkey")
		value := CreateRandomData()
		metadata := []byte("ttl:40;")

		g.It("should create a put instance of Request", func() {
			req := builder.CreatePutCommand(key, value, metadata)

			g.Assert(reflect.TypeOf(req).String()).Equal("*spotcache.Request")
			g.Assert(len(req.Id)).Equal(26)
			g.Assert(len(req.Session)).Equal(12)
		})

		g.It("should create an instance of RequestBuilder")
	})
}
