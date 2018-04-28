package main

import (
	"go-gin-api/controllers"

	"github.com/gin-gonic/gin"
)

// NewRouter Generate New Router
func NewRouter() *gin.Engine {
	router := gin.New()

	// Gin Middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// Controller
	ping := new(controllers.PingController)

	// Route List
	router.GET("/ping", ping.Ping)
	router.GET("/ping.json", ping.PingJSON)

	return router
}
