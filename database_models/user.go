package database_models

import (
	"main/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id uint64 `json:"id" gorm:"primaryKey"`
}

func init() {
	database.DB.AutoMigrate(&User{})
}

func Create(user *[]User) error {
	return database.DB.Create(user).Error
}

func Query(ids []int) ([]User, error) {
	var users []User
	if err := database.DB.Find(&users, ids).Error; err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func Update(ids []int, values map[string]interface{}) error {
	return database.DB.Model(User{}).Where("id IN ?", ids).Updates(values).Error
}

func Delete(ids []int) error {
	return database.DB.Model(User{}).Where("id IN ?", ids).Delete(&User{}).Error
}
