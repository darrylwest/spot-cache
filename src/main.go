/**
 * main
 *
 * @author darryl.west <darryl.west@raincitysoftware.com>
 * @created 2017-03-10 14:13:11
 */

package main

import (
    "fmt"
    "spotcache"
)

func main() {
    cfg := spotcache.ParseArgs()
    log := spotcache.CreateLogger(cfg)

    log.Info("Starting socket service started: %v\n", cfg.ToMap())
    err := spotcache.StartService(cfg)

    if err != nil {
        fmt.Println("error starting service: ", err)
        panic(err)
    }

}
