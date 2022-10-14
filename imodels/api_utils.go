package models

import (
	"github.com/gin-gonic/gin"
)

func RouteAll(router *gin.Engine) {
	if router == nil {
		panic("models::RouteAll: nil router")
	}
	// TODO REST routing
}
