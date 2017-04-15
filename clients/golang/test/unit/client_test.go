//
// ClientTest
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-15 10:05:33
//

package unit

import (
    "fmt"
	"spotclient"
	"testing"
    "time"

	. "github.com/franela/goblin"
)

func TestClient(t *testing.T) {
	g := Goblin(t)

	g.Describe("Client", func() {
		cfg := spotclient.NewConfigForEnvironment("test")

		g.It("should create a client struct", func() {
			client := spotclient.NewSpotClient(cfg)
            now := time.Now()

			g.Assert(client.Sess).Equal("")
            fmt.Printf("%v\n", *client)
            g.Assert(client.CreateTime.Hour()).Equal(now.Hour())
            g.Assert(client.CreateTime.Minute()).Equal(now.Minute())
		})
	})
}
