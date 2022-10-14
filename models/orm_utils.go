package models

import (
	"main/orm"
)

func OrmCreate[T IModels](objs []T) error {
	return orm.DB.Create(objs).Error
}

func OrmQuery[T IModels](where map[string]interface{}, order string) ([]T, error) {
	var objs []T
	tx := orm.DB
	if len(where) != 0 {
		tx = tx.Where(where)
	}
	if len(order) != 0 {
		tx = tx.Order(order)
	}
	if err := tx.Find(&objs).Error; err != nil {
		return nil, err
	} else {
		return objs, nil
	}
}

func OrmUpdate[T IModels](where map[string]interface{}, values map[string]interface{}) error {
	var obj T
	return orm.DB.Model(&obj).Where(where).Updates(values).Error
}

func OrmDelete[T IModels](where map[string]interface{}) error {
	var obj T
	return orm.DB.Where(where).Delete(&obj).Error
}
