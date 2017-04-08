//
// CommandTest
//
// @author darryl west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 13:56:46
//

package unit

import (
	// "fmt"
	"spotcache"
	"testing"

	. "github.com/franela/goblin"
)

func TestCommand(t *testing.T) {
	g := Goblin(t)

	var (
		ok   = []byte(spotcache.ResponseOk)
		fail = []byte(spotcache.ResponseFail)
		yes  = []byte(spotcache.ResponseTrue)
		// no = []byte(spotcache.ResponseFalse)
		pong       = []byte(spotcache.ResponsePong)
		knownValue = []byte("this is my test value")
		cache      *spotcache.Cache
	)

	g.Describe("Command", func() {
		cfg := spotcache.NewConfigForEnvironment("test")
		var session spotcache.SessionType
		copy(session[:12], []byte(spotcache.CreateSessionID()))
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
			key := []byte("MyTestKey")
			value := []byte("My Test Value with a specific length")
			metadata := []byte("expire:60;")
			request := builder.CreatePutRequest(key, value, metadata)

			g.Assert(request.Op).Equal(spotcache.PUT)
			bytes, err := request.ToBytes()

			g.Assert(err).Equal(nil)
			g.Assert(len(bytes) > 40)

			req, err := spotcache.RequestFromBytes(bytes)

			// fmt.Printf("req: %s\n", req)

			// TODO : move this to request tests
			g.Assert(err).Equal(nil)
			g.Assert(req.ID).Equal(request.ID)
			g.Assert(req.Session).Equal(request.Session)
			g.Assert(req.Op).Equal(request.Op)
			g.Assert(req.MetaSize).Equal(request.MetaSize)
			g.Assert(req.KeySize).Equal(request.KeySize)
			g.Assert(req.DataSize).Equal(request.DataSize)

			g.Assert(req.Metadata).Equal(request.Metadata)
			g.Assert(req.Key).Equal(request.Key)
			g.Assert(req.Value).Equal(request.Value)

			g.Assert(req).Equal(request)
		})

		g.It("should execute a put command", func() {
			id := spotcache.CreateCommandID()
			op := spotcache.PUT
			key := []byte("mykey")

			cmd := spotcache.CreateCommand(id, op, key, knownValue)

			err := cmd.Exec()
			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(ok)
		})

		g.It("should parse a get command")
		g.It("should execute a get command", func() {
			id := spotcache.CreateCommandID()
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
			id := spotcache.CreateCommandID()
			op := spotcache.HAS
			key := []byte("mykey")

			cmd := spotcache.CreateCommand(id, op, key, nil)

			err := cmd.Exec()
			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(yes)
		})

		g.It("should parse a del command")
		g.It("should execute a del command")

		g.It("should parse a ping command", func() {
			req := []byte("")
			cmd, err := spotcache.ParseRequest(req)

			g.Assert(err == nil).IsTrue()
			err = cmd.Exec()
			g.Assert(err == nil).IsTrue()

		})
		g.It("should execute a ping command", func() {
			id := spotcache.CreateCommandID()
			op := spotcache.PING
			cmd := spotcache.CreateCommand(id, op, nil, nil)
			err := cmd.Exec()

			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(pong)
		})

		g.It("should parse a stat command")
		g.It("should execute a stat command", func() {
			id := spotcache.CreateCommandID()
			op := spotcache.STATUS
			cmd := spotcache.CreateCommand(id, op, nil, nil)
			err := cmd.Exec()

			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(ok)
		})

		g.It("should parse a shutdown command")
		g.It("should execute a shutdown command", func() {
			id := spotcache.CreateCommandID()
			op := spotcache.SHUTDOWN
			cmd := spotcache.CreateCommand(id, op, nil, nil)
			err := cmd.Exec()

			g.Assert(err).Equal(nil)
			g.Assert(cmd.Resp).Equal(fail)
		})

		g.It("should reject an unknown command", func() {
			id := spotcache.CreateCommandID()
			op := spotcache.CommandOp(250)
			cmd := spotcache.CreateCommand(id, op, nil, nil)
			err := cmd.Exec()

			g.Assert(err != nil).IsTrue("error should not be nil")
			g.Assert(cmd.Resp).Equal(fail)
		})
	})
}
