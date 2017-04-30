// logger
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 10:57:53

package spotcache

import (
	"fmt"
	"path"

	"github.com/darrylwest/cassava-logger/logger"
)

var log *logger.Logger

// CreateLogger create a new logger based on config
func CreateLogger(cfg *Config) *logger.Logger {
	if log == nil {
		filename := path.Join(cfg.Logpath, cfg.Logname)
		handler, err := logger.NewRotatingDayHandler(filename)

		if err != nil {
            fmt.Printf("%s\n", err)
			panic("logger could not be created...")
		}

		fmt.Printf("created rolling log file: %s\n", filename)
		log = logger.NewLogger(handler)
	}

	return log
}
