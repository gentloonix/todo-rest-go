package orm

import (
	"main/models"
)

func OrmCreate[T models.IModels](objs []T) error {
	return db.Create(objs).Error
}

func OrmQuery[T models.IModels](where map[string]interface{}, order string) ([]T, error) {
	var objs []T
	tx := db
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

func OrmUpdate[T models.IModels](where map[string]interface{}, values map[string]interface{}) error {
	var obj T
	return db.Model(&obj).Where(where).Updates(values).Error
}

func OrmDelete[T models.IModels](where map[string]interface{}) error {
	var obj T
	return db.Where(where).Delete(&obj).Error
}
