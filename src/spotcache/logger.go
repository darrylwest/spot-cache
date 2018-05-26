// logger
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-11 10:57:53

package spotcache

import (
	"fmt"
	"os"
	"path"

	"github.com/darrylwest/cassava-logger/logger"
)

var log *logger.Logger

// CreateLogger create a new logger based on config
func CreateLogger(cfg *Config) *logger.Logger {
	if log == nil {
		if err := CreateLogFolder(cfg.Logpath); err != nil {
			panic(err)
		}

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

// CreateLogFolder - create the log folder if it doesn't already exist
func CreateLogFolder(logpath string) error {
	if _, err := os.Stat(logpath); err == nil {
		return nil
	}

	fmt.Printf("create log folder %s\n", logpath)
	return os.MkdirAll(logpath, 0755)
}
