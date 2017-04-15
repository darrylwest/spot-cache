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

type Config struct {
    Home string
    Env string
    Logpath string
    Longnme string
    Port int
    Timeout int64
}

func NewDefaultConfig() *Config {
    cfg := new(Config)

    home := os.Getenv("SPOTCACHE_HOME")
    if home == "" {
        home = path.Join(os.Getenv("HOME"), ".spotcache")
    }

    cfg.Home = home
    cfg.Env = "production"
    cfg.Logpath = path.Join(home, "logs")
    cfg.Logname = Sprintf("spotclient-%d", os.Getpid())
    cfg.Port = 3001
    cfg.Timeout = int64(10 * 60)

    return cfg
}

