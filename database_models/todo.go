package database_models

type TODO struct {
	ORMWrapper
	Id          uint64 `json:"id" gorm:"primaryKey"`
	UserId      User   `json:"user_id" gorm:"foreignKey:Id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	DueDate     int64  `json:"due_date,omitempty"`
	Status      bool   `json:"status,omitempty"`
}
