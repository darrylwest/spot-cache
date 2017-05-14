//
// spotcache cli
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-14 08:33:42
//

package main

import (
	"fmt"
	"spotclient"
)

func main() {
	cfg := spotclient.ParseArgs()
    client := spotclient.NewSpotClient(cfg)

    fmt.Printf("%v\n", client)
}
