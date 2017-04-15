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

	"spotcache"
)

func main() {
	cfg := spotcache.ParseArgs()
	log := spotcache.CreateLogger(cfg)
	defer log.Close()
	// TODO : do this in config and write the pid to the pid file...
	pid := os.Getpid()

	service := spotcache.NewCacheService(cfg)
	service.InitializeCache(cfg)

	ss, err := service.CreateListener()
	if err != nil {
		panic(err)
	}

	log.Info("Starting socket service started: %v, pid: %d\n", cfg, pid)

	go service.ListenAndServe(ss)
	defer service.Shutdown()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigchan
	log.Info("signal caught: %v", sig)

	fmt.Println("signal caught:", sig)
	fmt.Println("stopped...")
}
