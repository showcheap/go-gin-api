package controllers

import (
	"fmt"
	"go-gin-api/db"
	"go-gin-api/models"

	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct{}

// Index ...
func (u UserController) Index(c *gin.Context) {
	db := db.Instance()
	var users []models.User

	if err := db.Find(&users).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, users)
	}

}

// Detail ...
func (u UserController) Detail(c *gin.Context) {
	db := db.Instance()
	id := c.Params.ByName("id")

	var user models.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"message": "User Not Found!"})
		fmt.Println(err)
		return
	}

	c.JSON(200, user)
}

// Create ...
func (u UserController) Create(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	db := db.Instance()
	db.Create(&user)

	c.JSON(200, user)
}

// Delete ...
func (u UserController) Delete(c *gin.Context) {
	db := db.Instance()

	id := c.Params.ByName("id")
	var user models.User

	if err := db.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"message": "User Not Found!"})
		fmt.Println(err)
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(404, gin.H{"message": "Unable to delete user"})
		fmt.Println(err)
		return
	}

	c.JSON(200, gin.H{"message": "User deleted!"})

}
