package main

import (
	"io"
	"net"
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
	c, err := net.Dial("unix", "/src/tmp/echo.sock")
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
