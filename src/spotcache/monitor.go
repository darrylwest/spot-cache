//
// Monitor - create the monitor socket; listen for requests; process and return responses; broadcast events
//
// @author darryl west <darryl.west@raincitysoftware.com>
// @created 2017-03-17 08:32:48
//

package spotcache

import (
	"fmt"
	"net"
	"os"
	"time"
)

type Monitor struct {
	Sockfile   string
	CreateDate time.Time
}

// simple request/response
type MonitorCommand struct {
	id   []byte
	op   []byte // shutdown, status, ping
	resp []byte
}

func NewMonitor(cfg *Config) Monitor {
    m := Monitor{}

    m.Sockfile = cfg.unixsock
	m.CreateDate = time.Now().UTC()

    return m
}

func (m *Monitor) OpenAndServe() {
	defer os.Remove(m.Sockfile)
	ss, err := net.Listen("unix", m.Sockfile)

	if err != nil {
		panic(err)
	}

	defer ss.Close()

    // set the start listen date/time
	m.CreateDate = time.Now().UTC()

	for {
		conn, err := ss.Accept()
		if err != nil {
			log.Error("Unix socket connection fail: %v", err)
			break
		}

		go m.ClientHandler(conn)
	}

	log.Info("closing monitor service")
}

// handle a new monitor client
func (m *Monitor) ClientHandler(conn net.Conn) {
	buf := make([]byte, 512)
	defer conn.Close()

	sess, err := StartClientSession(conn)
	if err != nil {
		log.Info("session error, aborting...")
		return
	}

	log.Info("started a monitor client session: %s", sess)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Error("client monitor error: %v", err)
			break
		}

		// parse and respond to command...
		request := buf[:n]
		fmt.Printf("rcvd: %s", request)

		_, err = conn.Write([]byte(fmt.Sprintf(fmt.Sprintf("echo %s", request))))
		if err != nil {
			log.Error("error responding to client request, closing session: %s", sess)
			break
		}
	}

	log.Info("closing session: %s", sess)
}
