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
	Session    string
}


// NewSpotClient - create the client
func NewSpotClient(cfg *Config) *SpotClient {
	client := &SpotClient{}

	client.cfg = cfg
	client.CreateTime = time.Now()

	return client
}

// Exec - run the command
func (client SpotClient) Exec() error {
    var err error

    fmt.Printf("exec %v\n", client.cfg.Args) 

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
