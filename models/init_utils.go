package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AutoMigrate initializes all possible database tables
func AutoMigrate(db *gorm.DB) {
	if db == nil {
		panic("models::AutoMigrate: nil db")
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&TODO{}); err != nil {
		panic(err)
	}
}

// RouteAll routes all possible paths
func RouteAll(router *gin.Engine) {
	if router == nil {
		panic("models::RouteAll: nil router")
	}
	RouteTODO(router)
}
