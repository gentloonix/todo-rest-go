package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TODO struct {
	gorm.Model
	Id          uint64 `json:"id" gorm:"primaryKey"`
	UserId      User   `json:"user_id" gorm:"foreignKey:Id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DueDate     int64  `json:"due_date,omitempty"`
	Status      bool   `json:"status,omitempty"`
}

func RouteTODO(router *gin.Engine) {
	router.GET("/todo/", ApiGet[TODO])
	router.POST("/todo/", ApiPost[TODO])
	router.PUT("/todo/", ApiPut[TODO])
	router.DELETE("/todo/", ApiDelete[TODO])
}
