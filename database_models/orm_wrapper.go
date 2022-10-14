package database_models

import (
	"main/database"

	"gorm.io/gorm"
)

type ORMWrapper struct {
	gorm.Model
}

func Create(objs []ORMWrapper) error {
	return database.DB.Create(objs).Error
}

func Query(values map[string]interface{}) ([]ORMWrapper, error) {
	var objs []ORMWrapper
	if err := database.DB.Where(values).Find(&objs).Error; err != nil {
		return nil, err
	} else {
		return objs, nil
	}
}

func Update(ids []int, values map[string]interface{}) error {
	return database.DB.Where("id IN ?", ids).Updates(values).Error
}

func Delete(ids []int) error {
	return database.DB.Where("id IN ?", ids).Delete(&ORMWrapper{}).Error
}
