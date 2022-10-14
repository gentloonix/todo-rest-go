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
	// Close channel
	Close chan struct{} = make(chan struct{}, 1)
	// router pointer
	router *gin.Engine = nil
)

// Initialize server submodule as daemon
func Initialize(wg *sync.WaitGroup) {
	if wg == nil {
		log.Println("api::Initialize: nil wg")
		return
	}
	if router != nil {
		log.Println("api::Initialize: ignoring")
		return
	}

	go func() {
		defer wg.Done()

		defer func() {
			log.Println("api: exit")
		}()

		// release mode
		gin.SetMode(gin.ReleaseMode)

		// based on https://github.com/gin-gonic/gin#manually
		router = gin.Default()
		defer func() {
			router = nil
		}()

		routeAll(router)

		srv := &http.Server{
			Addr:    ADDR,
			Handler: router,
		}

		// Initializing the server in a goroutine so that
		// it won't block the graceful shutdown handling below
		go func() {
			if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
				log.Println("api:", err)
			}
		}()

		log.Println("api: running")
		<-Close

		// The context is used to inform the server it has 30 seconds to finish
		// the request it is currently handling
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Println("api: forceful shutdown:", err)
		}
	}()
}
