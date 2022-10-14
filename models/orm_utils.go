package models

import "main/orm"

type ORMModel interface {
	TODO
	User
}

func Create[T ORMModel](objs []T) error {
	return orm.DB.Create(objs).Error
}

func Query[T ORMModel](values map[string]interface{}) ([]T, error) {
	var objs []T
	if err := orm.DB.Where(values).Find(&objs).Error; err != nil {
		return nil, err
	} else {
		return objs, nil
	}
}

func Update[T ORMModel](ids []int, values map[string]interface{}) error {
	var obj T
	return orm.DB.Model(&obj).Where("id IN ?", ids).Updates(values).Error
}

func Delete[T ORMModel](ids []int) error {
	var obj T
	return orm.DB.Where("id IN ?", ids).Delete(&obj).Error
}
