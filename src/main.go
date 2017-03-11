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
    /*
    err := spotcache.StartServer(cfg)

    if err != nil {
        fmt.Println("error starting service: ", err)
        panic(err)
    }
    */

    fmt.Printf("Socket service started: %v\n", cfg.ToMap())
}
