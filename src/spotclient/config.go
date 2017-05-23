//
// Config
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-04-15 08:50:25
//

package spotclient

import (
	"flag"
	"fmt"
	"os"
	"path"
)

// Config - client configuration struct
type Config struct {
	Env     string
	Host    string
	Port    int
	Timeout int64
	Args    []string
}

// NewDefaultConfig - create a new config using the standard defaults
func NewDefaultConfig() *Config {
	cfg := new(Config)

	cfg.Env = "production"
	cfg.Host = "localhost"
	cfg.Port = 19501
	cfg.Timeout = int64(10 * 60)

	return cfg
}

// ParseArgs - parse the command line args to set host, port, env etc
func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")

	env := flag.String("env", dflt.Env, "set the environment, defaults to "+dflt.Env)
	host := flag.String("host", dflt.Host, "set the server's host, defaults to "+dflt.Host)
	port := flag.Int("port", dflt.Port, fmt.Sprintf("set the listening port, defaults to %d", dflt.Port))

	flag.Parse()

	fmt.Printf("%s Version %s\n", path.Base(os.Args[0]), Version())

	if *vers == true {
		os.Exit(0)
	}

	cfg := new(Config)

	cfg.Env = *env
	cfg.Host = *host
	cfg.Port = *port

	// copy over the res of the args
	cfg.Args = flag.Args()

	cfg.Timeout = dflt.Timeout

	return cfg
}

// NewConfigForEnvironment - create new config for the given environment
func NewConfigForEnvironment(env string) *Config {
	cfg := NewDefaultConfig()

	cfg.Env = env

	return cfg
}
