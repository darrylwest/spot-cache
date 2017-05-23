//
// SpotClient
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-15 08:40:52
//

package spotclient

import (
	"fmt"
	"net"
	"time"

	"github.com/darrylwest/spot-cache/spotcache"
)

// SpotClient - client struct
type SpotClient struct {
	CreateTime time.Time
	cfg        *Config
	Session    spotcache.SessionType
	builder    *spotcache.RequestBuilder
}

// NewSpotClient - create the client
func NewSpotClient(cfg *Config) *SpotClient {
	client := &SpotClient{}

	client.cfg = cfg
	client.CreateTime = time.Now()

	return client
}

func (client *SpotClient) getSession(conn net.Conn) spotcache.SessionType {
	buf := make([]byte, 32)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	copy(client.Session[:], buf[:n])

	return client.Session
}

// SendPing - sends a basic ping to the server
func (client *SpotClient) SendPing(conn net.Conn, count int, interval time.Duration) error {
	fmt.Printf("send a ping %d requests...\n", count)

	buf := make([]byte, 128)
	for i := 0; i < count; i++ {
		if i > 0 {
			time.Sleep(interval)
		}

		request := client.builder.CreatePingRequest()
		bytes, _ := request.ToBytes()

		if _, err := conn.Write(bytes); err != nil {
			fmt.Println("lost connection...")
			return err
		}

		fmt.Printf("ping request  #%d: %v\n", (i + 1), request)

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("lost connection..")
			return err
		}

		resp, _ := spotcache.ResponseFromBytes(buf[:n])
		fmt.Printf("ping response #%d: ID:%s SS: %s data:%s\n", (i + 1), resp.ID, resp.Session, string(resp.Data))
	}

	return nil
}

// Exec - run the command
func (client SpotClient) Exec() error {
	var err error

	fmt.Printf("exec %v\n", client.cfg.Args)

	// implement ping first...
	conn, _ := client.Connect()
	defer conn.Close()

	sess := client.getSession(conn)
	fmt.Printf("session: %s\n", sess)

	client.builder = spotcache.NewRequestBuilder(sess)

	// now send a ping
	client.SendPing(conn, 1e6, time.Second*10)

	return err
}

// Connect - return the connection
func (client *SpotClient) Connect() (net.Conn, error) {
	host := fmt.Sprintf("%s:%d", client.cfg.Host, client.cfg.Port)

	fmt.Printf("dialing %s\n", host)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		panic(err)
	}

	return conn, err
}

// Send a command request
func (client *SpotClient) Send(request *spotcache.Request) error {
	return nil
}
