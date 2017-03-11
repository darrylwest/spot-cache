package spotcache

import (
	"flag"
	"fmt"
	"os"
	"path"
)

type Config  struct {
	env           string // defaults to production
	logpath       string
	logname       string
    dbpath        string
	baseport      int
	shutdownPort  int
	timeout       int64
}

func (c *Config) GetShutdownPort() int {
	return c.shutdownPort
}

func (c *Config) ToMap() map[string]interface{} {
	hash := make(map[string]interface{})

	hash["env"] = c.env
	hash["logpath"] = c.logpath
	hash["logname"] = c.logname

    hash["dbpath"] = c.dbpath
	hash["baseport"] = c.baseport
	hash["shutdownPort"] = c.shutdownPort
	hash["timeout"] = c.timeout

	return hash
}

func NewDefaultConfig() *Config {
	cfg := new(Config)

	cfg.env = "production"
	cfg.logpath = path.Join(os.Getenv("HOME"), "logs")
	cfg.logname = "spotcache"

    cfg.dbpath = "cachedb"

	cfg.baseport = 3001
	cfg.shutdownPort = 3009

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

	env := flag.String("env", dflt.env, "set the environment, defaults to "+dflt.env)

	baseport := flag.Int("baseport", dflt.baseport, "set the server's base port number (e.g., 3001)...")
	shutdownPort := flag.Int("shutdownPort", dflt.shutdownPort, "set the service shutdown port")

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

	cfg.env = *env
	cfg.logpath = *logpath
	cfg.logname = *logname

	cfg.dbpath = *dbpath

	cfg.baseport = *baseport
	cfg.shutdownPort = *shutdownPort
	cfg.timeout = *timeout

	return cfg
}

func IsProduction(env string) bool {
    return env == "production"
}
