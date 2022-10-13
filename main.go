package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"main/database"
	_ "main/database"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		r := gin.Default()

		r.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "hello world"})
		})

		r.Run()

		log.Println("shut down")
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	database.Initialize(&wg)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	database.Close <- struct{}{}

	wg.Wait()
}
