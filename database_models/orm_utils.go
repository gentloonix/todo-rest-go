package database_models

import (
	"main/database"

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

func Create[T any](objs []T) error {
	return database.DB.Create(objs).Error
}

func Query[T any](values map[string]interface{}) ([]T, error) {
	var objs []T
	if err := database.DB.Where(values).Find(&objs).Error; err != nil {
		return nil, err
	} else {
		return objs, nil
	}
}

func Update[T any](ids []int, values map[string]interface{}) error {
	var obj T
	return database.DB.Model(&obj).Where("id IN ?", ids).Updates(values).Error
}

func Delete[T any](ids []int) error {
	var obj T
	return database.DB.Where("id IN ?", ids).Delete(&obj).Error
}
