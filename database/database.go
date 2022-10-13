package database

import (
	"log"
	"main/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Close chan struct{} = make(chan struct{}, 1)
	DB    *gorm.DB      = nil
)

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

			DB = nil
		}()

		if err := db.AutoMigrate(&models.User{}); err != nil {
			panic(err)
		}
		if err := db.AutoMigrate(&models.TODO{}); err != nil {
			panic(err)
		}

		DB = db

		<-Close
		// TODO waitGroup for synchronized exiting
	}()
}
