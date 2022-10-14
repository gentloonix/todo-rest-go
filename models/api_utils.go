package models

import (
	"log"

	"github.com/gin-gonic/gin"
)

func ApiGet[T IModels](c *gin.Context) {
	log.Printf("%#v\n", c.Request.URL.Query())
}

func ApiPost[T IModels](c *gin.Context) {
	log.Printf("%#v\n", c.Request.URL.Query())
}

func ApiPut[T IModels](c *gin.Context) {
	log.Printf("%#v\n", c.Request.URL.Query())
}

func ApiDelete[T IModels](c *gin.Context) {
	log.Printf("%#v\n", c.Request.URL.Query())
}
