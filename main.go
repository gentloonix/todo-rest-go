package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"main/api"
	"main/orm"
)

func main() {
	// Async submodules initialize
	log.Println("submodules initializing")

	var wg sync.WaitGroup

	wg.Add(1)
	orm.Initialize(&wg, "database.db")
	wg.Add(1)
	api.Initialize(&wg)

	// Await SIGINT / SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Await async submodules exit
	log.Println("submodules exiting")

	orm.Close <- struct{}{}
	api.Close <- struct{}{}

	wg.Wait()
}
