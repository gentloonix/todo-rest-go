package database_models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id uint64 `json:"id" gorm:"primaryKey"`
}
