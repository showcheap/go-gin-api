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

	router.GET("/panic", func(c *gin.Context) {
		panic("tes")
	})

	// Route Group API
	api := router.Group("/api")
	{
		api.GET("/", ping.Welcome)

	}

	apiUser := api.Group("/user")
	{
		apiUser.GET("/", userC.Index)
		apiUser.POST("/", userC.Create)
		apiUser.GET("/:id", userC.Detail)
		apiUser.DELETE("/:id", userC.Delete)
	}

	return router
}
