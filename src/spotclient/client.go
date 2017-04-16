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
	Sess       string
}

// NewSpotClient - create the client
func NewSpotClient(cfg *Config) *SpotClient {
	client := &SpotClient{}

	client.cfg = cfg
	client.CreateTime = time.Now()

	return client
}

// Connect - return the connection
func (client *SpotClient) Connect() (net.Conn, error) {
	host := fmt.Sprintf("%s:%d", client.cfg.Host, client.cfg.Port)

	// fmt.Println(host);

	conn, err := net.Dial("tcp", host)

	return conn, err
}

// Send a command request
func (client *SpotClient) Send(request *spotcache.Request) error {
	return nil
}
