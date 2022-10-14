package database_models

import (
	"main/database"

	"gorm.io/gorm"
)

func Create[T *gorm.Model](objs []T) error {
	return database.DB.Create(objs).Error
}

func Query[T *gorm.Model](values map[string]interface{}) ([]T, error) {
	var objs []T
	if err := database.DB.Where(values).Find(&objs).Error; err != nil {
		return nil, err
	} else {
		return objs, nil
	}
}

func Update[T *gorm.Model](ids []int, values map[string]interface{}) error {
	var obj T
	return database.DB.Model(&obj).Where("id IN ?", ids).Updates(values).Error
}

func Delete[T *gorm.Model](ids []int) error {
	var obj T
	return database.DB.Where("id IN ?", ids).Delete(&obj).Error
}
