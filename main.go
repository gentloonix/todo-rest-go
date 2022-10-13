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
	var wg sync.WaitGroup

	wg.Add(1)
	database.Initialize(&wg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	log.Println("exiting...")

	database.Close <- struct{}{}

	wg.Wait()
}
