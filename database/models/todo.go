package models

import "gorm.io/gorm"

type TODO struct {
	gorm.Model
	Id          uint64 `json:"id" gorm:"primaryKey"`
	UserId      User   `json:"user_id" gorm:"foreignKey:Id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     int64  `json:"due_date"`
	Status      bool   `json:"status"`
}
