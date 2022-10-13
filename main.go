package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"main/api"
	"main/database"
)

func main() {
	// Async submodules initialize
	var wg sync.WaitGroup

	wg.Add(1)
	database.Initialize(&wg)
	wg.Add(1)
	api.Initialize(&wg)

	// Await SIGINT / SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Await async submodules exit
	log.Println("exiting...")

	database.Close <- struct{}{}
	api.Close <- struct{}{}

	wg.Wait()
}
