package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"main/database"
)

func main() {
	// Async submodules initialize
	var wg sync.WaitGroup

	wg.Add(1)
	database.Initialize(&wg)

	// Await SIGINT / SIGTERM
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Await async submodules exit
	log.Println("exiting...")

	database.Close <- struct{}{}

	wg.Wait()
}
