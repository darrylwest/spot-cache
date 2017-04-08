package main

import (
	"fmt"
	"net"
	"os"
	"time"

	_ "spotcache"
)

func getSession(conn net.Conn) string {
	buf := make([]byte, 250)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	return string(buf[:n])
}

func main() {
	id := time.Now().UnixNano()

	port := 3001
	host := fmt.Sprintf(":%d", port)
	fmt.Println("dailing: ", host)

	conn, err := net.Dial("tcp", host)
	if err != nil {
		fmt.Println("error connecting to host ", host)
		os.Exit(1)
	}

	defer conn.Close()
	count := 0

	key := fmt.Sprintf("client:%d", id)

	sess := getSession(conn)
	fmt.Printf("my session: %s\n", sess)

	buf := make([]byte, 2048)
	for {
		count++
		text := fmt.Sprintf("put:%d %s 'my value %d'", count, key, time.Now().Unix())
		_, err := fmt.Fprintf(conn, text)
		if err != nil {
			fmt.Println("lost connection...")
			return
		}

		fmt.Printf("request: %s\n", text)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("lost connection...")
			return
		}

		fmt.Printf("resp: %s\n", buf[:n])

		time.Sleep(time.Second)
	}
}
