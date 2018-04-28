package controllers

import (
	"github.com/gin-gonic/gin"
)

// UserController ...
type UserController struct{}

// Index ...
func (u UserController) Index(c *gin.Context) {
	c.String(200, "User index")
}

// Detail ...
func (u UserController) Detail(c *gin.Context) {
	c.String(200, "User detail")
}

// Create ...
func (u UserController) Create(c *gin.Context) {
	c.String(200, "Create User")
}

// Delete ...
func (u UserController) Delete(c *gin.Context) {
	c.String(200, "Delete User")
}
