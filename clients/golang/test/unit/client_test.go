//
// ClientTest
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-15 10:05:33
//

package unit

import (
	"spotclient"
	"testing"

	. "github.com/franela/goblin"
)

func TestClient(t *testing.T) {
	g := Goblin(t)

	g.Describe("Client", func() {
		cfg := spotclient.NewConfigForEnvironment("test")

		g.It("should create a client struct", func() {
			client := spotclient.NewSpotClient(cfg)

			g.Assert(client.Sess).Equal("")
		})
	})
}
