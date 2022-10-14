package api

import (
	"main/models"

	"github.com/gin-gonic/gin"
)

// routeAll routes all possible paths
func routeAll(router *gin.Engine) {
	if router == nil {
		panic("models::RouteAll: nil router")
	}
	routeUser(router)
	routeTODO(router)
}

func routeUser(router *gin.Engine) {
	router.GET("/user/", ApiGet[models.User])
	router.POST("/user/", ApiPost[models.User])
	router.PUT("/user/", ApiPut[models.User])
	router.DELETE("/user/", ApiDelete[models.User])
}

func routeTODO(router *gin.Engine) {
	router.GET("/todo/", ApiGet[models.TODO])
	router.POST("/todo/", ApiPost[models.TODO])
	router.PUT("/todo/", ApiPut[models.TODO])
	router.DELETE("/todo/", ApiDelete[models.TODO])
}
