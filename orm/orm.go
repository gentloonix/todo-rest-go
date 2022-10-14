package orm

import (
	"log"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	// Close close channel
	Close chan struct{} = make(chan struct{}, 1)
	// DB pointer to database
	DB *gorm.DB = nil
)

// Initialize initializes Database ORM (Object-Relational Mapping) submodule as daemon;
// stops on Close signal, decrements wg on exit;
// autoMigrate initializes all possible database tables;
func Initialize(wg *sync.WaitGroup, dsn string, autoMigrate func(*gorm.DB)) {
	if wg == nil {
		log.Println("orm::Initialize: nil wg")
		return
	}
	if DB != nil {
		log.Println("orm::Initialize: ignoring")
		return
	}

	go func() {
		defer wg.Done()

		defer func() {
			log.Println("orm: exit")
		}()

		var err error
		DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			QueryFields: true,
		})
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

		autoMigrate(DB)

		log.Println("orm: running")
		<-Close
	}()
}
