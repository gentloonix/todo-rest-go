package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id     uint64 `json:"id" gorm:"primaryKey"`
	Email  string `json:"email"`
	Secret string `json:"secret"` // argon2id
}
