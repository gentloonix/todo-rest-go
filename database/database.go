package database

import (
	"log"
	"main/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	go func() {
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

		if err := db.AutoMigrate(&models.User{}); err != nil {
			panic(err)
		}
		if err := db.AutoMigrate(&models.TODO{}); err != nil {
			panic(err)
		}

		DB = db

		// TODO block until main() begins exiting
		// TODO waitGroup for synchronized exiting
	}()
}
