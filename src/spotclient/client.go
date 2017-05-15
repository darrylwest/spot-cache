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
func (client *SpotClient) SendPing(builder *spotcache.RequestBuilder, conn net.Conn) error {
    buf := make([]byte, 128)
    fmt.Println("send a ping request")
    request := builder.CreatePingRequest()
    bytes, _ := request.ToBytes()

    if _, err := conn.Write(bytes); err != nil {
        fmt.Println("lost connection...");
        return err
    }

    fmt.Printf("request: %v\n", request)

    n, err := conn.Read(buf)
    if err != nil {
        fmt.Println("lost connection..");
        return err
    }

    resp, _ := spotcache.ResponseFromBytes(buf[:n])
    fmt.Printf("ping response: ID: %s SS: %s data:%s\n", resp.ID, resp.Session, string(resp.Data))

    return nil
}

// Exec - run the command
func (client SpotClient) Exec() error {
    var err error

    fmt.Printf("exec %v\n", client.cfg.Args) 

    // implement ping first...
    conn, _ := client.Connect()
    defer conn.Close();

    sess := client.getSession(conn)
    fmt.Printf("session: %s\n", sess);

    // now send a ping
    builder := spotcache.NewRequestBuilder(sess)
    client.SendPing(builder, conn)

    return err
}

// Connect - return the connection
func (client *SpotClient) Connect() (net.Conn, error) {
	host := fmt.Sprintf("%s:%d", client.cfg.Host, client.cfg.Port)

	fmt.Printf("dialing %s\n", host);
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
