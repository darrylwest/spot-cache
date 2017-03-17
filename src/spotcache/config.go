//
// config  - application specification and CLI parsing
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 13:56:46

package spotcache

import (
	"flag"
	"fmt"
	"os"
	"path"
)

type Config struct {
	home     string // defaults to user's home
	env      string // defaults to production
	logpath  string
	logname  string
	dbpath   string
	baseport int
	unixsock string
	timeout  int64
}

func (c *Config) GetUnixSock() string {
	return c.unixsock
}

func (c *Config) ToMap() map[string]interface{} {
	hash := make(map[string]interface{})

	hash["home"] = c.home
	hash["env"] = c.env
	hash["logpath"] = c.logpath
	hash["logname"] = c.logname

	hash["dbpath"] = c.dbpath
	hash["baseport"] = c.baseport
	hash["unixsock"] = c.unixsock
	hash["timeout"] = c.timeout

	return hash
}

func NewDefaultConfig() *Config {
	cfg := new(Config)

	home := os.Getenv("SPOTCACHE_HOME")
	if home == "" {
		home = path.Join(os.Getenv("HOME"), ".spotcache")
	}

	cfg.home = home

	cfg.env = "production"
	cfg.logpath = path.Join(home, "logs")
	cfg.logname = "spotcache"

	cfg.dbpath = path.Join(home, "cachedb")

	cfg.baseport = 3001
	cfg.unixsock = path.Join(home, "spot.sock")

	cfg.timeout = int64(10 * 60) // seconds in unix time

	return cfg
}

func NewConfigForEnvironment(env string) *Config {
	cfg := NewDefaultConfig()

	cfg.env = env

	if !IsProduction(env) {
		cfg.logname = env + "-spotcache"
	}

	return cfg
}

func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")

	home := flag.String("home", dflt.home, "set the run-time home folder, defaults to "+os.Getenv("HOME"))
	env := flag.String("env", dflt.env, "set the environment, defaults to "+dflt.env)

	baseport := flag.Int("baseport", dflt.baseport, "set the server's base port number (e.g., 3001)...")
	unixsock := flag.String("unixsock", dflt.unixsock, "set the service status/shutdown socket")

	logpath := flag.String("logpath", dflt.logpath, "set the log directory")
	logname := flag.String("logname", dflt.logname, "set the name of the rolling log file")

	dbpath := flag.String("dbpath", dflt.dbpath, "set the database directory")
	timeout := flag.Int64("timeout", dflt.timeout, "set the timeout in seconds")

	flag.Parse()

	fmt.Printf("%s Version: %s\n", path.Base(os.Args[0]), Version())

	if *vers == true {
		os.Exit(0)
	}

	cfg := new(Config)

	cfg.home = *home
	cfg.env = *env
	cfg.logpath = *logpath
	cfg.logname = *logname

	cfg.dbpath = *dbpath

	cfg.baseport = *baseport
	cfg.unixsock = *unixsock
	cfg.timeout = *timeout

	return cfg
}

func IsProduction(env string) bool {
	return env == "production"
}
