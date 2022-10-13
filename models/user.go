package models

import (
	"fmt"
	"main/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id uint64 `json:"id" gorm:"primaryKey"`
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	database.DB.Create(&user)
	c.JSON(200, user)
}

func ReadUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func ReadUsers(c *gin.Context) {
	var users []User
	if err := database.DB.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}
}

func UpdateUser(c *gin.Context) {
	var user User
	id := c.Params.ByName("id")
	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&user)
	database.DB.Save(&user)
	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	d := database.DB.Where("id = ?", id).Delete(&user)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
