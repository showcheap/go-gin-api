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
	userC := new(controllers.UserController)

	// Route List
	router.GET("/", ping.Welcome)
	router.GET("/ping", ping.Ping)
	router.GET("/ping.json", ping.PingJSON)

	// Route Group API
	api := router.Group("/api")
	{
		api.GET("/", ping.Welcome)

	}

	apiUser := api.Group("/user")
	{
		apiUser.GET("/", userC.Index)
	}

	return router
}
