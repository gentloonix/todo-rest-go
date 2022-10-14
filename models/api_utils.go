package models

import (
	"github.com/gin-gonic/gin"
)

func RouteAll(g *gin.Engine) {
	if g == nil {
		panic("models::RouteAll: nil g")
	}
}
