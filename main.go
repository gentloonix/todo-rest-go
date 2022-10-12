package main

import (
	"log"

	_ "github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id uint64
}

type TODO struct {
	gorm.Model
	UserId      uint64
	Name        string
	Description string
	DueDate     uint64
	Status      bool
}

func main() {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Println(err)
		}
	}()

	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&TODO{}); err != nil {
		panic(err)
	}
}
