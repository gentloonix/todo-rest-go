package database_models

import (
	"main/database"

	"gorm.io/gorm"
)

type TODO struct {
	gorm.Model
	Id          uint64 `json:"id" gorm:"primaryKey"`
	UserId      User   `json:"user_id" gorm:"foreignKey:Id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DueDate     int64  `json:"due_date,omitempty"`
	Status      bool   `json:"status,omitempty"`
}

func init() {
	database.DB.AutoMigrate(&TODO{})
}
