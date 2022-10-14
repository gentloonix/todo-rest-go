package database_models

import (
	"main/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id uint64 `json:"id" gorm:"primaryKey"`
}

func Create(users []User) error {
	return database.DB.Create(users).Error
}

func Query(values map[string]interface{}) ([]User, error) {
	var users []User
	if err := database.DB.Where(values).Find(&users).Error; err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func Update(ids []int, values map[string]interface{}) error {
	return database.DB.Where("id IN ?", ids).Updates(values).Error
}

func Delete(ids []int) error {
	return database.DB.Where("id IN ?", ids).Delete(&User{}).Error
}
