package spotcache_test

import (
	"spotcache"
	"testing"
	// "fmt"

	. "github.com/franela/goblin"
)

func TestCommand(t *testing.T) {
	g := Goblin(t)

	var (
		ok = []byte(spotcache.RESP_OK)
		// fail = []byte(spotcache.RESP_FAIL)
		yes = []byte(spotcache.RESP_TRUE)
		// no = []byte(spotcache.RESP_FALSE)
		pong       = []byte(spotcache.RESP_PONG)
		knownValue = []byte("this is my test value")
	)

	g.Describe("Command", func() {
		g.Before(func() {
			conf := spotcache.NewConfigForEnvironment("test")
			spotcache.CreateLogger(conf)
			spotcache.Opendb(conf)
		})

		g.After(func() {
			spotcache.Closedb()
		})

		g.It("should parse a put command")
		g.It("should execute a put command", func() {
			id := []byte("01BB01P6QMY0DJB7V412A29TJB")
			op := []byte("put")
			key := []byte("mykey")

			cmd := spotcache.CreateCommand(id, op, key, knownValue)

			err := cmd.Exec()
			g.Assert(err).Equal(nil)
			g.Assert(cmd.GetResp()).Equal(ok)
		})

		g.It("should parse a get command")
		g.It("should execute a get command", func() {
			id := []byte("01BB01P6QMY0DJB7V412A29TJB")
			op := []byte("get")
			key := []byte("mykey")

			cmd := spotcache.CreateCommand(id, op, key, nil)

			err := cmd.Exec()
			g.Assert(err).Equal(nil)
			// fmt.Println(cmd.ToString())
			// now check for the response
			g.Assert(cmd.GetResp()).Equal(knownValue)
		})

		g.It("should parse a has command")
		g.It("should execute a has command", func() {
			id := []byte("01BB01P6QMY0DJB7V412A29TJB")
			op := []byte("has")
			key := []byte("mykey")

			cmd := spotcache.CreateCommand(id, op, key, nil)

			err := cmd.Exec()
			g.Assert(err).Equal(nil)
			g.Assert(cmd.GetResp()).Equal(yes)
		})

		g.It("should parse a del command")
		g.It("should execute a del command")

		g.It("should parse a ping  command")
		g.It("should execute a ping command", func() {
			id := []byte("01BB01P6QMY0DJB7V412A29TJB")
			op := []byte("ping")
			cmd := spotcache.CreateCommand(id, op, nil, nil)
			err := cmd.Exec()

			g.Assert(err).Equal(nil)
			g.Assert(cmd.GetResp()).Equal(pong)
		})

		g.It("should parse a stat command")
		g.It("should execute a stat command")

		g.It("should parse a halt command")
		g.It("should execute a halt command")

	})
}
