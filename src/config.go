package webserver

import (
	"flag"
	"fmt"
	"os"
	"path"
)

type Config  struct {
	env           string // defaults to production
	webroot       string
	logpath       string
	logname       string
	baseport      int
	shutdownPort  int
	serverCount   int
	indexTemplate string
	appTemplate   string
	static        string
	timeout       int64
}

func (c *Config) GetShutdownPort() int {
	return c.shutdownPort
}

func (c *Config) ToMap() map[string]interface{} {
	hash := make(map[string]interface{})

	hash["env"] = c.env
	hash["webroot"] = c.webroot
	hash["logpath"] = c.logpath
	hash["logname"] = c.logname

	hash["baseport"] = c.baseport
	hash["shutdownPort"] = c.shutdownPort

	hash["serverCount"] = c.serverCount
	hash["indexTemplate"] = c.indexTemplate
	hash["appTemplate"] = c.appTemplate
	hash["static"] = c.static
	hash["timeout"] = c.timeout

	return hash
}

func NewDefaultConfig() *Config {
	ctx := new(Config)

	ctx.env = "production"
	ctx.webroot = "public"
	ctx.static = "static"
	ctx.logpath = path.Join(os.Getenv("HOME"), "logs")
	ctx.logname = "webserver"

	ctx.baseport = 3001
	ctx.shutdownPort = 3009

	ctx.serverCount = 2
	ctx.timeout = int64(10 * 60) // seconds in unix time

	return ctx
}

func NewConfigForEnvironment(env string) *Config {
	ctx := NewDefaultConfig()

	ctx.env = env

	if !IsProduction(env) {
		ctx.logname = env + "-webserver"
	}

	return ctx
}

func ParseArgs() *Config {
	dflt := NewDefaultConfig()

	vers := flag.Bool("version", false, "show the version and exit")

	env := flag.String("env", dflt.env, "set the environment, defaults to "+dflt.env)

	baseport := flag.Int("baseport", dflt.baseport, "set the server's base port number (e.g., 3001)...")
	serverCount := flag.Int("serverCount", dflt.serverCount, "set the number of server/listeners")
	shutdownPort := flag.Int("shutdownPort", dflt.shutdownPort, "set the service shutdown port")

	webroot := flag.String("webroot", dflt.webroot, "set the path to the server's static file directory, (e.g., 'public')...")
	static := flag.String("static", dflt.static, "set the hidden static file service directory")
	logpath := flag.String("logpath", dflt.logpath, "set the log directory")
	logname := flag.String("logname", dflt.logname, "set the name of the rolling log file")

	index := flag.String("index", dflt.indexTemplate, "set the index template filename")
	app := flag.String("app", dflt.appTemplate, "set the application template filename")

	timeout := flag.Int64("timeout", dflt.timeout, "set the timeout in seconds")

	flag.Parse()

	fmt.Printf("%s Version: %s\n", path.Base(os.Args[0]), Version())

	if *vers == true {
		os.Exit(0)
	}

	ctx := new(Config)

	ctx.env = *env
	ctx.webroot = *webroot
	ctx.static = *static
	ctx.logpath = *logpath
	ctx.logname = *logname

	ctx.baseport = *baseport
	ctx.shutdownPort = *shutdownPort
	ctx.serverCount = *serverCount
	ctx.indexTemplate = *index
	ctx.appTemplate = *app
	ctx.timeout = *timeout

	return ctx
}
