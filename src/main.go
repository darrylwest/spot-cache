// main
//
// @author darryl.west <darryl.west@raincitysoftware.com>
// @created 2017-03-10 14:13:11

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"spotcache"
)

func main() {
	cfg := spotcache.ParseArgs()
	log := spotcache.CreateLogger(cfg)
	service := spotcache.NewCacheService(cfg)

	sigchan := make(chan os.Signal, 1)
	stop := make(chan bool)

	go func() {
		sig := <-sigchan
		log.Info("recived signal %v", sig)
		stop <- true

		time.Sleep(time.Duration(50 * time.Millisecond))
		log.Info("shutdown complete...")
		time.Sleep(time.Duration(1 * time.Second))
	}()

	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	pid := os.Getpid()
	log.Info("Starting socket service started: %v, pid: %d\n", cfg.ToMap(), pid)
	fmt.Println("pid", pid)

	service.OpenAndServe(stop)

	time.Sleep(time.Duration(100 * time.Millisecond))
}
