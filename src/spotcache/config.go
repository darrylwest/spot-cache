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

// Config the config structure
type Config struct {
	Home     string // defaults to user's home
	Env      string // defaults to production
	Logpath  string
	Logname  string
	Dbpath   string
	Baseport int
	Unixsock string
	Timeout  int64
}

// NewDefaultConfig default settings
func NewDefaultConfig() *Config {
	cfg := new(Config)

	home := os.Getenv("SPOTCACHE_HOME")
	if home == "" {
		home = path.Join(os.Getenv("HOME"), ".spotcache")
	}

	cfg.Home = home

	cfg.Env = "production"
	cfg.Logpath = path.Join(home, "logs")
	cfg.Logname = "spotcache"

	cfg.Dbpath = path.Join(home, "cachedb")

	cfg.Baseport = 3001
	cfg.Unixsock = path.Join(home, "spot.sock")

	cfg.Timeout = int64(10 * 60) // seconds in unix time

	return cfg
}

// NewConfigForEnvironment configure for a specific environment
func NewConfigForEnvironment(env string) *Config {
	cfg := NewDefaultConfig()

	cfg.Env = env

	if !IsProduction(env) {
		cfg.Logname = env + "-spotcache"
	}

	return cfg
}

// ParseArgs parse the command line args
func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")

	home := flag.String("home", dflt.Home, "set the run-time home folder, defaults to "+os.Getenv("HOME"))
	env := flag.String("env", dflt.Env, "set the environment, defaults to "+dflt.Env)

	baseport := flag.Int("baseport", dflt.Baseport, "set the server's base port number (e.g., 3001)...")
	unixsock := flag.String("unixsock", dflt.Unixsock, "set the service status/shutdown socket")

	logpath := flag.String("logpath", dflt.Logpath, "set the log directory")
	logname := flag.String("logname", dflt.Logname, "set the name of the rolling log file")

	dbpath := flag.String("dbpath", dflt.Dbpath, "set the database directory")
	timeout := flag.Int64("timeout", dflt.Timeout, "set the timeout in seconds")

	flag.Parse()

	fmt.Printf("%s Version: %s\n", path.Base(os.Args[0]), Version())

	if *vers == true {
		os.Exit(0)
	}

	cfg := new(Config)

	cfg.Home = *home
	cfg.Env = *env
	cfg.Logpath = *logpath
	cfg.Logname = *logname

	cfg.Dbpath = *dbpath

	cfg.Baseport = *baseport
	cfg.Unixsock = *unixsock
	cfg.Timeout = *timeout

	return cfg
}

// IsProduction return true if the current env is production
func IsProduction(env string) bool {
	return env == "production"
}
