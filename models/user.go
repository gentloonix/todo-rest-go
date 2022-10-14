package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id     uint64 `json:"id" gorm:"primaryKey"`
	Email  string `json:"email"`
	Secret string `json:"secret"` // argon2id
}

func RouteUser(router *gin.Engine) {
	router.GET("/user/", ApiGet[User])
	router.POST("/user/", ApiPost[User])
	router.PUT("/user/", ApiPut[User])
	router.DELETE("/user/", ApiDelete[User])
}
