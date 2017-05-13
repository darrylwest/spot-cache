//
// writer client - write domain key/value pairs to cache; read k/v list from a refernce database
//
// 1) connect; 2) create n k/v records; 3) write all k/v recors to cache; 4) report timing and exit.
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-05-13 13:30:10
//

package main

import (
	"fmt"
	"net"
	// "os"
	"time"

	"github.com/darrylwest/go-unique/unique"
	"spotcache"
)

var maxCount = 1000

// User a sample domain object
type User struct {
	id          string
	dateCreated time.Time
	lastUpdated time.Time
	version     int64
	username    string
	group       string
	status      string
}

func getSession(conn net.Conn) spotcache.SessionType {
	buf := make([]byte, 64)
	n, err := conn.Read(buf)
	if err != nil {
		panic(err)
	}

	var ss spotcache.SessionType

	copy(ss[:], buf[:n])

	return ss
}

func createUser(group string) User {
	now := time.Now()

	user := User{}
	user.id = unique.CreateULID()
	user.dateCreated = now
	user.lastUpdated = now

	if rbytes, err := unique.RandomBytes(16); err == nil {
		user.username = fmt.Sprintf("%x", rbytes)
	}

	user.group = group

	user.status = "new"

	return user
}

func createUserList() []User {
	group := unique.CreateTSID()
	list := make([]User, maxCount)

	for i := 0; i < maxCount; i++ {
		list[i] = createUser(group)
	}

	return list[:]
}

func parseArgs() *spotcache.Config {
	cfg := spotcache.NewDefaultConfig()

	return cfg
}

func main() {
	cfg := parseArgs()

	fmt.Println("max count: ", maxCount)
	fmt.Println(cfg)

	list := createUserList()

	for _, user := range list {
		fmt.Println(user)
	}

	// connect and get the session

}
