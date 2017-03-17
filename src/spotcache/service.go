// service - the TCP interface
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 11:42:53

package spotcache

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

type CacheService struct {
	ClientCount int
	CreateDate  time.Time
	Port        int
	DbPath      string
}

func NewCacheService(cfg *Config) CacheService {
	s := CacheService{}

	s.Port = cfg.baseport
	s.DbPath = cfg.dbpath

	return s
}

// open the cache database and start the main socket service; block forever...
func (s *CacheService) OpenAndServe(stop <-chan bool) {
	OpenDb(s.DbPath)
	defer CloseDb()

	// OpenSocketService
	host := fmt.Sprintf(":%d", s.Port)
	ss, err := net.Listen("tcp", host)

	if err != nil {
		log.Error("error creating connection: %v", err)
		return
	}

	defer ss.Close()
	log.Info("listinging on port: %s", host)

	go func() {
		for {
			conn, err := ss.Accept()
			if err != nil {
				log.Error("error on accept: %v", err)
				break
			}

			go s.OpenClientHandler(conn)

			// should probably shove the clint conn into a map
		}
	}()

	// wait for the stop message
	<-stop
}

// handle client requests as long as they stay connected
func (s *CacheService) OpenClientHandler(conn net.Conn) {
	buf := make([]byte, 8192)
	defer conn.Close()

	sess, err := StartClientSession(conn)
	if err != nil {
		log.Info("session error, aboring...")
		return
	}

	log.Info("session started: %s", sess)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Warn("client connection error, connection lost...")
			break
		}

		log.Info("REQ: %s", buf[:n])

		resp := fmt.Sprintf("%s", buf[:n])
		log.Info("RSP: %s", resp)

		fmt.Fprintf(conn, resp)
	}
}

// create a client session id and send to the new client (move to sock utils?)
func StartClientSession(conn net.Conn) (string, error) {
	sess := strconv.FormatInt(time.Now().UTC().UnixNano(), 36)

	if _, err := fmt.Fprintf(conn, sess); err != nil {
		log.Error("session create error: %v", err)
		return sess, err
	} else {
		log.Info("started session: %s", sess)
	}

	return sess, nil
}
