//
// SpotClient
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-15 08:40:52
//

package spotclient

import (
    "fmt"
    // spotcache "github.com/darrylwest/spot-cache/"
)

type SpotClient struct {
}

func NewClient() *SpotClient {
    client := &SpotClient{}

    return client
}
