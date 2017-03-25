//
// CommandTest
//
// @author darryl west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 13:56:46
//

package test

import (
	"spotcache"
	"testing"

	. "github.com/franela/goblin"
)

func TestCommand(t *testing.T) {
	g := Goblin(t)

	var (
		ok   = []byte(spotcache.RESP_OK)
		fail = []byte(spotcache.RESP_FAIL)
		yes  = []byte(spotcache.RESP_TRUE)
		// no = []byte(spotcache.RESP_FALSE)
		pong       = []byte(spotcache.RESP_PONG)
		knownValue = []byte("this is my test value")
		cache      *spotcache.Cache
	)

	g.Describe("Command", func() {
		cfg := spotcache.NewConfigForEnvironment("test")
		session := []byte("test1234")
		builder := spotcache.NewRequestBuilder(session)

		g.Before(func() {
			spotcache.CreateLogger(cfg)
			cache = spotcache.NewCache(cfg)

			cache.Open()
		})

		g.After(func() {
			cache.Close()
		})

		g.It("should parse a put command", func() {
			key := []byte("mytestkey")
			value := CreateRandomData()
			metadata := []byte("ttl:60;")
			request := builder.CreatePutCommand(key, value, metadata)

			g.Assert(string(request.Op)).Equal("pu")
			// cmd := spotcache.ParseRequest(ToByteArray(request))
			// fmt.Println( request.ToByteArray() );
		})

		g.It("should execute a put command", func() {
			id := CreateCommandId()
			op := spotcache.PUT
			key := []byte("mykey")

			cmd := spotcache.CreateCommand(id, op, key, knownValue)

			err := cmd.Exec()
			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(ok)
		})

		g.It("should parse a get command")
		g.It("should execute a get command", func() {
			id := CreateCommandId()
			op := spotcache.GET
			key := []byte("mykey")

			cmd := spotcache.CreateCommand(id, op, key, nil)

			err := cmd.Exec()
			g.Assert(err).Equal(nil)
			// fmt.Println(cmd.ToString())
			// now check for the response
			g.Assert(cmd.Resp).Equal(knownValue)
		})

		g.It("should parse a has command")
		g.It("should execute a has command", func() {
			id := CreateCommandId()
			op := spotcache.HAS
			key := []byte("mykey")

			cmd := spotcache.CreateCommand(id, op, key, nil)

			err := cmd.Exec()
			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(yes)
		})

		g.It("should parse a del command")
		g.It("should execute a del command")

		g.It("should parse a ping  command", func() {
			req := []byte("")
			cmd, err := spotcache.ParseRequest(req)

			g.Assert(err == nil).IsTrue()
			err = cmd.Exec()

		})
		g.It("should execute a ping command", func() {
			id := CreateCommandId()
			op := spotcache.PING
			cmd := spotcache.CreateCommand(id, op, nil, nil)
			err := cmd.Exec()

			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(pong)
		})

		g.It("should parse a stat command")
		g.It("should execute a stat command", func() {
			id := CreateCommandId()
			op := spotcache.STATUS
			cmd := spotcache.CreateCommand(id, op, nil, nil)
			err := cmd.Exec()

			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(ok)
		})

		g.It("should parse a shutdown command")
		g.It("should execute a shutdown command", func() {
			id := CreateCommandId()
			op := spotcache.SHUTDOWN
			cmd := spotcache.CreateCommand(id, op, nil, nil)
			err := cmd.Exec()

			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(fail)
		})

		g.It("should reject an unknown command", func() {
			id := CreateCommandId()
			op := spotcache.CommandOp(250)
			cmd := spotcache.CreateCommand(id, op, nil, nil)
			err := cmd.Exec()

			g.Assert(err != nil).IsTrue("error should not be nil")
			g.Assert(cmd.Resp).Equal(fail)
		})
	})
}
