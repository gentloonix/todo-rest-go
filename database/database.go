package database

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Close chan struct{} = make(chan struct{}, 1)
	DB    *gorm.DB      = nil
)

func Initialize(wg *sync.WaitGroup) {
	if wg == nil {
		log.Println("database::Initialize: nil wg")
		return
	}
	if DB != nil {
		log.Println("database::Initialize: ignoring")
		return
	}

	go func() {
		defer wg.Done()

		defer func() {
			log.Println("database: exit")
		}()

		DB, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		sqlDB, err := DB.DB()
		if err != nil {
			panic(err)
		}
		defer func() {
			if err := sqlDB.Close(); err != nil {
				log.Println(err)
			}

			DB = nil
		}()

		log.Println("database: running")
		<-Close
	}()
}
