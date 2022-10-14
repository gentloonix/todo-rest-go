package models

import (
	"main/models"
)

type ORMModel interface {
	models.TODO
	models.User
}
