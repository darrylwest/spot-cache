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
			g.Assert(client.CreateTime.Hour()).Equal(now.Hour())
			g.Assert(client.CreateTime.Minute()).Equal(now.Minute())
		})

		g.It("should reject a connection to a non-listening host/port", func() {
			client := spotclient.NewSpotClient(cfg)

			conn, err := client.Connect()

			g.Assert(conn).Equal(nil)
			g.Assert(err != nil).Equal(true)
			g.Assert(fmt.Sprintf("%s", err)).Equal("dial tcp [::1]:3001: getsockopt: connection refused")
			// fmt.Printf("%s\n", err)
		})

		g.It("should return an error if send is invoked with a non-open connection", func() {
			client := spotclient.NewSpotClient(cfg)

			g.Assert(client != nil).IsTrue()
		})
	})
}
