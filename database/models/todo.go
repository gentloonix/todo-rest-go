package models

import "gorm.io/gorm"

type TODO struct {
	gorm.Model
	UserId      uint64
	Name        string
	Description string
	DueDate     uint64
	Status      bool
}
