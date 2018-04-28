package controllers

import "github.com/gin-gonic/gin"

// PingController ...
type PingController struct{}

// Ping Handler
func (p PingController) Ping(c *gin.Context) {
	c.String(200, "Pong!")
}

// PingJSON Ping Handler in JSON format
func (p PingController) PingJSON(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Pong!"})
}

// Welcome ...
func (p PingController) Welcome(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Welcome"})
}
