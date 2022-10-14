package database_models

import "gorm.io/gorm"

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&TODO{})
	db.AutoMigrate(&User{})
}
