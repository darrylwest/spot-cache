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

func sendPing(builder *spotcache.RequestBuilder, conn net.Conn) error {
	buf := make([]byte, 128)
	fmt.Println("send a ping request")
	request := builder.CreatePingRequest()
	bytes, _ := request.ToBytes()

	_, err := conn.Write(bytes)
	if err != nil {
		fmt.Println("lost connection...")
		return err
	}

	fmt.Printf("request: %v\n", request)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("lost connection...")
		return err
	}

	resp, _ := spotcache.ResponseFromBytes(buf[:n])
	fmt.Printf("ping response: ID: %s SS: %s data:%s\n", resp.ID, resp.Session, string(resp.Data))

	return nil
}

func parseArgs() *spotcache.Config {
	cfg := spotcache.NewDefaultConfig()

	return cfg
}

func main() {
	cfg := parseArgs()

	host := fmt.Sprintf(":%d", cfg.Baseport)
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
	messageCount := 0
	t0 := time.Now().UnixNano()
	for {
		count++
		req := builder.CreateGetRequest([]byte("MyTestKey"), nil)
		bytes, _ := req.ToBytes()

		_, err := conn.Write(bytes)
		if err != nil {
			fmt.Println("lost connection...")
			return
		}

		// fmt.Printf("request : %v\n", request)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("lost connection...")
			return
		}

		resp, err := spotcache.ResponseFromBytes(buf[:n])
		if err != nil {
			fmt.Println("error %v", err)
			return
		}

		if resp.ID != req.ID || resp.Session != resp.Session || resp.DataSize == 0 {
			fmt.Printf("response error: ID:%s Session:%s Op: %d data:%s\n", resp.ID, resp.Session, resp.Op, string(resp.Data))
		}

		messageCount++

		if count%5000 == 0 {
			t1 := time.Now().UnixNano()
			fmt.Printf("Total messages sent/received: %d %f millis\n", messageCount, (float64(t1)-float64(t0))/1e6)
			time.Sleep(time.Second)
			err = sendPing(builder, conn)
			if err != nil {
				fmt.Println("ping died...")
				return
			}

			t0 = time.Now().UnixNano()
		}
	}
}
