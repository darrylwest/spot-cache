//
// monitor client
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-08 09:55:41
//
package main

import (
	"io"
	"net"
	"os"
	"path"
	"strings"
	"time"
)

func reader(r io.Reader) {
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf[:])
		if err != nil {
			return
		}
		println("Client got:", string(buf[0:n]))
	}
}

func main() {
	home := os.Getenv("SPOTCACHE_HOME")
	if home == "" {
		home = path.Join(os.Getenv("HOME"), ".spotcache")
	}

	sock := path.Join(home, "spot.sock")

	c, err := net.Dial("unix", sock)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	go reader(c)
	for {
		s := []string{"ping", time.Now().UTC().Format("2006-01-02T15:04:05.000Z")}

		_, err := c.Write([]byte(strings.Join(s, " ")))
		if err != nil {
			println(err.Error())
			break
		}

		time.Sleep(1e9)
	}
}
