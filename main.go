package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "Pong!")
	})

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome!")
	})

	r.GET("/ping.json", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Pong!"})
	})

	r.Run(":8081")
}
