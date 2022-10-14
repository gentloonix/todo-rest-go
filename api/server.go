package api

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	Close  chan struct{} = make(chan struct{}, 1)
	Router *gin.Engine   = nil
)

// Initialize initializes server submodule as daemon
// Stops on Close signal
// routeAll routes all possible paths
func Initialize(wg *sync.WaitGroup, routeAll func(*gin.Engine)) {
	if wg == nil {
		log.Println("api::server::Initialize: nil wg")
		return
	}
	if Router != nil {
		log.Println("api::server::Initialize: ignoring")
		return
	}

	go func() {
		defer wg.Done()

		defer func() {
			log.Println("api::server: exit")
		}()

		// release mode
		gin.SetMode(gin.ReleaseMode)

		// based on https://github.com/gin-gonic/gin#manually
		Router = gin.Default()
		defer func() {
			Router = nil
		}()

		routeAll(Router)

		srv := &http.Server{
			Addr:    ":8080",
			Handler: Router,
		}

		// Initializing the server in a goroutine so that
		// it won't block the graceful shutdown handling below
		go func() {
			if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
				log.Println("api::server:", err)
			}
		}()

		log.Println("api::server: running")
		<-Close

		// The context is used to inform the server it has 30 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Println("api::server: forceful shutdown:", err)
		}
	}()
}
