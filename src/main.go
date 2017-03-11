/**
 * main
 *
 * @author darryl.west <darryl.west@raincitysoftware.com>
 * @created 2017-03-10 14:13:11
 */

package main

import (
    "fmt"
    "service"
)

func main() {
    cfg := config.ParseArgs()
    err := service.StartServer(cfg)

    if err != nil {
        fmt.Println("error starting service: ", err)
        panic(err)
    }

    fmt.Printf("Socket service started: $v\n", cfg);
}
