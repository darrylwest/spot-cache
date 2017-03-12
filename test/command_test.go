package spotcache_test

import (
	"testing"
	"spotcache"
    "fmt"

	. "github.com/franela/goblin"
)

func TestCommand(t *testing.T) {
	g := Goblin(t)

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
            value := []byte("this is my test value")

            cmd := spotcache.CreateCommand(id, op, key, value)
            
            err := cmd.Exec()
            g.Assert(err).Equal(nil)
            g.Assert(cmd.GetResp()).Equal(spotcache.RESP_OK)
        })

		g.It("should parse a get command")
		g.It("should execute a get command", func() {
            id := []byte("01BB01P6QMY0DJB7V412A29TJB")
            op := []byte("get")
            key := []byte("mykey")

            cmd := spotcache.CreateCommand(id, op, key, nil)
            
            err := cmd.Exec()
            g.Assert(err).Equal(nil)
            fmt.Println(cmd.ToString())
        })

		g.It("should parse a has command")
		g.It("should execute a has command")

		g.It("should parse a del command")
		g.It("should execute a del command")

		g.It("should parse a stat command")
		g.It("should execute a stat command")

		g.It("should parse a halt command")
		g.It("should execute a halt command")

	})
}
