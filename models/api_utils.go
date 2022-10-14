package models

import (
	"github.com/gin-gonic/gin"
)

// RouteAll routes all possible paths
func RouteAll(router *gin.Engine) {
	if router == nil {
		panic("models::RouteAll: nil router")
	}
	RouteTODO(router)
}

func RESTGet[T IModels](c *gin.Context) {
}

func RESTPost[T IModels](c *gin.Context) {

}

func RESTPut[T IModels](c *gin.Context) {

}

func RESTDelete[T IModels](c *gin.Context) {

}
