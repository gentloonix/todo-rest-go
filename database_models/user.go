package database_models

type User struct {
	ORMWrapper
	Id uint64 `json:"id" gorm:"primaryKey"`
}
