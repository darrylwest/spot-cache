// service - the TCP interface
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 11:42:53

package spotcache

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type CacheService struct {
	ClientCount int
	CreateDate  time.Time
	Port        int
	DbPath      string
	Timeout     time.Duration
	stopchan    chan bool
	waitGroup   *sync.WaitGroup
}

func NewCacheService(cfg *Config) CacheService {
	s := CacheService{}

	s.Port = cfg.Baseport
	s.DbPath = cfg.Dbpath
	s.Timeout = 1e9

	s.stopchan = make(chan bool, 1)
	s.waitGroup = &sync.WaitGroup{}

	return s
}

func (s *CacheService) CreateListener() (*net.TCPListener, error) {
	host := fmt.Sprintf("127.0.0.1:%d", s.Port)
	laddr, err := net.ResolveTCPAddr("tcp", host)

	ss, err := net.ListenTCP("tcp", laddr)

	return ss, err
}

// open the cache database and start the main socket service; block forever...
func (s *CacheService) ListenAndServe(ss *net.TCPListener) {
	s.CreateDate = time.Now().UTC()

	// TODO: open and pass the database into service...
	OpenDb(s.DbPath)
	defer CloseDb()

	defer ss.Close()
	log.Info("listinging on port: %v", ss.Addr())

	defer s.waitGroup.Done()

	for {
		select {
		case <-s.stopchan:
			log.Error("stopping...")
			ss.Close()
			return
		default:
		}

		to := time.Now().Add(s.Timeout)

		log.Debug("%v", to)
		ss.SetDeadline(to)
		conn, err := ss.Accept()

		if err != nil {
			if opErr, ok := err.(*net.OpError); ok && opErr.Timeout() {
				continue
			}

			log.Error("error on accept: %v", err)
			continue
		}

		// should probably shove the clint conn into a conection array/slice
		log.Info("connection on %v", conn.RemoteAddr())
		s.waitGroup.Add(1)
		go s.OpenClientHandler(conn)
	}
}

func (s *CacheService) Shutdown() {
	close(s.stopchan)
	s.waitGroup.Wait()
}

// handle client requests as long as they stay connected
func (s *CacheService) OpenClientHandler(conn net.Conn) {
	buf := make([]byte, 8192)
	defer conn.Close()
	defer s.waitGroup.Done()

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
