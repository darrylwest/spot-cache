//
// Config
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-15 08:50:25
//

package spotclient

import (
	// "flag"
	"fmt"
	"os"
	"path"
)

// Config - client configuration struct
type Config struct {
	Home    string
	Env     string
	Logpath string
	Logname string
	Host    string
	Port    int
	Timeout int64
}

// NewDefaultConfig - create a new config using the standard defaults
func NewDefaultConfig() *Config {
	cfg := new(Config)

	home := os.Getenv("SPOTCACHE_HOME")
	if home == "" {
		home = path.Join(os.Getenv("HOME"), ".spotcache")
	}

	cfg.Home = home
	cfg.Env = "production"
	cfg.Logpath = path.Join(home, "logs")
	cfg.Logname = fmt.Sprintf("spotclient-%d", os.Getpid())
	cfg.Host = "localhost"
	cfg.Port = 3001
	cfg.Timeout = int64(10 * 60)

	return cfg
}

//func ParseArgs() *Config {
//}

// NewConfigForEnvironment - create new config for the given environment
func NewConfigForEnvironment(env string) *Config {
	cfg := NewDefaultConfig()

	cfg.Env = env

	return cfg
}
