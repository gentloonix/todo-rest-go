package orm

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// Close channel
	Close chan struct{} = make(chan struct{}, 1)
	// db pointer to database
	db *gorm.DB = nil
)

// Initialize initializes Database ORM (Object-Relational Mapping) submodule as daemon
func Initialize(wg *sync.WaitGroup) {
	if wg == nil {
		log.Println("orm::Initialize: nil wg")
		return
	}
	if db != nil {
		log.Println("orm::Initialize: ignoring")
		return
	}

	go func() {
		defer wg.Done()

		defer func() {
			log.Println("orm: exit")
		}()

		var err error
		db, err = gorm.Open(sqlite.Open(DSN), &gorm.Config{
			QueryFields: true,
		})
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

			db = nil
		}()

		autoMigrate(db)

		log.Println("orm: running")
		<-Close
	}()
}
