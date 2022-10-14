package database_models

import (
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	if db == nil {
		panic("database_models::AutoMigrate: nil db")
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&TODO{}); err != nil {
		panic(err)
	}
}
