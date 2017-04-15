//
// SpotClient
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-15 08:40:52
//

package spotclient

/*
import (
    // "fmt"
    // spotcache "github.com/darrylwest/spot-cache/"
)
*/

// SpotClient - client struct
type SpotClient struct {
	Sess string
	cfg  *Config
}

// NewSpotClient - create the client
func NewSpotClient(cfg *Config) *SpotClient {
	client := &SpotClient{}

	client.cfg = cfg

	return client
}
