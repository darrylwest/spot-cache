/**
 * logger
 *
 * @author darryl.west <darryl.west@raincitysoftware.com>
 * @created 2017-03-11 10:57:53
 */

package spotcache

import (
	"fmt"
	"path"

	"github.com/darrylwest/cassava-logger/logger"
)

var log *logger.Logger

func CreateLogger(cfg *Config) *logger.Logger {
	if log == nil {
		filename := path.Join(cfg.logpath, cfg.logname)
		handler, err := logger.NewRotatingDayHandler(filename)

		if err != nil {
			panic("logger could not be created...")
		}

		fmt.Printf("created rolling log file: %s\n", filename)
		log = logger.NewLogger(handler)
	}

	return log
}
