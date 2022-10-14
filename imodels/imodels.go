package imodels

import (
	"main/models"
)

type IModels interface {
	models.TODO
	models.User
}
