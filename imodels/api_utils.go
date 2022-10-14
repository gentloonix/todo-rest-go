package models

import (
	"github.com/gin-gonic/gin"
)

// RouteAll routes all possible paths
func RouteAll(router *gin.Engine) {
	if router == nil {
		panic("models::RouteAll: nil router")
	}
	// TODO REST routing
}
