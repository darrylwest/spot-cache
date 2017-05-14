//
// Config - the client config
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-14 09:36:57
//

package clienttest

import (
	"fmt"
	"spotclient"
	"testing"

	. "github.com/franela/goblin"
)

func TestConfig(t *testing.T) {
	g := Goblin(t)

	g.Describe("ClientConfig", func() {

		g.It("should create a config struct", func() {
            cfg := spotclient.NewConfigForEnvironment("test")

			g.Assert(cfg != nil).IsTrue()

			fmt.Println(cfg)
		})

		// implement parse
	})
}

