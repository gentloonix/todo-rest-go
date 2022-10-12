package models

import "gorm.io/gorm"

type TODO struct {
	gorm.Model
	Id          uint64 `gorm:"primaryKey"`
	UserId      uint64
	Name        string
	Description string
	DueDate     uint64
	Status      bool
}
