package orm

import (
	"main/models"

	"gorm.io/gorm"
)

// autoMigrate initializes all possible database tables
func autoMigrate(db *gorm.DB) {
	if db == nil {
		panic("orm::autoMigrate: nil db")
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&models.TODO{}); err != nil {
		panic(err)
	}
}
