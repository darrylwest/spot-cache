//
// test client
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-08 09:55:56
//

package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"spotcache"
)

func getSession(conn net.Conn) spotcache.SessionType {
	buf := make([]byte, 64)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	var ss spotcache.SessionType

	copy(ss[:], buf[:n])
	return ss
}

func sendPing(conn net.Conn) {
	fmt.Println("send a ping request")
	request := builder.CreateGetRequest([]byte("MyTestKey"), nil)
	bytes, _ := request.ToBytes()

	_, err := conn.Write(bytes)
	if err != nil {
		fmt.Println("lost connection...")
		return
	}

	fmt.Printf("request: %v\n", request)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("lost connection...")
		return
	}

	fmt.Printf("resp: %s\n", buf[:n])
}

func main() {
	// id := time.Now().UnixNano()
	// key := fmt.Sprintf("client:%d", id)

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

	sess := getSession(conn)
	fmt.Printf("my session: %s\n", sess)

	builder := spotcache.NewRequestBuilder(sess)

	buf := make([]byte, 2048)
	for {
		count++
		request := builder.CreateGetRequest([]byte("MyTestKey"), nil)
		bytes, _ := request.ToBytes()

		_, err := conn.Write(bytes)
		if err != nil {
			fmt.Println("lost connection...")
			return
		}

		fmt.Printf("request: %v\n", request)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("lost connection...")
			return
		}

		fmt.Printf("resp: %s\n", buf[:n])

		time.Sleep(time.Second)
		sendPing()
	}
}
