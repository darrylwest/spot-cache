/**
 * service - the TCP interface
 *
 * @author darryl.west <darryl.west@raincitysoftware.com>
 * @created 2017-03-11 11:42:53
 */

package spotcache

import (
	"fmt"
	"net"
)

/**
 * open the cache database and start the main socket service; block forever...
 */
func StartService(cfg *Config) error {
    OpenDb(cfg)
    defer CloseDb()

	host := fmt.Sprintf(":%d", cfg.baseport)
	ss, err := net.Listen("tcp", host)

	if err != nil {
		log.Error("error creating connection: %v", err)
		return err
	}

	defer ss.Close()
	log.Info("listinging on port: %s", host)

	for {
		conn, err := ss.Accept()
		if err != nil {
			log.Error("error on accept: ", err.Error())
		}

		go handleClient(conn)
	}
}

/**
 * handle client requests as long as they stay connected
 */
func handleClient(conn net.Conn) {
	buf := make([]byte, 8192)
	defer conn.Close()

	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Warn("client connection error, connection lost...")
			return
		}

		log.Info("REQ: %s", buf[:n])

		resp := fmt.Sprintf("%s", buf[:n])
		log.Info("RSP: %s", resp)

		fmt.Fprintf(conn, resp)
	}
}
