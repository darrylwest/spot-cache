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

// CacheService - the main struct with client count, port, etc
type CacheService struct {
	ClientCount int
	CreateDate  time.Time
	Port        int
	DbPath      string
	Timeout     time.Duration
	stopchan    chan bool
	waitGroup   *sync.WaitGroup
}

var command *Commander
var cachedb *Cache

// NewCacheService - create a cache service
func NewCacheService(cfg *Config) CacheService {
	s := CacheService{}

	s.Port = cfg.Baseport
	s.DbPath = cfg.Dbpath
	s.Timeout = 1e9

	s.stopchan = make(chan bool, 1)
	s.waitGroup = &sync.WaitGroup{}

	return s
}

// InitializeCache configure the commander and cache database
func (s *CacheService) InitializeCache(cfg *Config) {
	cache := NewCache(cfg)
	command = NewCommander(cache)
}

// CreateListener create the listener for the specified address/port
func (s *CacheService) CreateListener() (*net.TCPListener, error) {
	host := fmt.Sprintf("127.0.0.1:%d", s.Port)
	laddr, err := net.ResolveTCPAddr("tcp", host)

	ss, err := net.ListenTCP("tcp", laddr)

	return ss, err
}

// ListenAndServe open the cache database and start the main socket service; block forever...
func (s *CacheService) ListenAndServe(ss *net.TCPListener) {
	s.CreateDate = time.Now().UTC()

	if command == nil || cache == nil {
		log.Error("must initialize cache prior to calling ListenAndServe...")
		panic("initialize error")
	}

	// open the cache db
	log.Info("open the cache database...")
	if err := cache.Open(); err != nil {
		log.Error("error opening cache, aborting...")
		panic(err)
	}

	defer cache.Close()
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

		// should probably shove the clint conn into a connection array/slice
		log.Info("connection on %v", conn.RemoteAddr())
		s.waitGroup.Add(1)
		go s.OpenClientHandler(conn)
	}
}

// Shutdown - stop everything
func (s *CacheService) Shutdown() {
	close(s.stopchan)
	s.waitGroup.Wait()
}

// OpenClientHandler handle client requests as long as they stay connected
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

		log.Info("req: %s", buf[:n])

		// read, decode and parse the request object
		req, err := RequestFromBytes(buf[:n])
		if err != nil {
			log.Error("%v", err)
			break
		}

		// create the command and execute it
		cmd := CreateCommand(req.ID, req.Op, req.Key, req.Value)
		err = cmd.Exec()
		if err != nil {
			log.Error("%s", err)
			break
		}

		log.Info("resp: %s\n", cmd.Resp)

		// TODO create a response object from request ID, Op, Key and Resp

		// return the response object to requester
		conn.Write(cmd.Resp)
	}

	log.Info("session %s closed...", sess)
}

// StartClientSession create a client session id and send to the new client (move to sock utils?)
func StartClientSession(conn net.Conn) (string, error) {
	sess := CreateSessionID()

	if _, err := fmt.Fprintf(conn, sess); err != nil {
		log.Error("session create error: %v", err)
		return sess, err
	}

	log.Info("started session: %s", sess)

	return sess, nil
}

// CreateSessionID returns a string of 12 chars
func CreateSessionID() string {
	sess := strconv.FormatInt(time.Now().UTC().UnixNano(), 36)
	if len(sess) == 12 {
		return sess
	} else if len(sess) < 12 {
		return (sess + "000")[:12]
	} else {
		return sess[:12]
	}
}
