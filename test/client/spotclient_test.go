//
// SpotClientTests
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-14 09:04:10
//

package clienttest

import (
	"fmt"
	"spotclient"
	// "spotcache"
	"testing"
	"time"

	. "github.com/franela/goblin"
)

func TestClient(t *testing.T) {
	g := Goblin(t)

	g.Describe("SpotClient", func() {
		cfg := spotclient.NewConfigForEnvironment("test")
		now := time.Now()

		g.It("should create a client struct", func() {

			client := spotclient.NewSpotClient(cfg)
			g.Assert(client != nil).IsTrue()
			g.Assert(client.CreateTime.After(now)).IsTrue()

			fmt.Sprintf("%v", client)
		})

		// implement ping, status, get, put/set, has, delete, backup
	})
}
