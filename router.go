package main

import (
	"go-gin-api/controllers"

	"github.com/gin-gonic/gin"
)

// NewRouter Generate New Router
func NewRouter() *gin.Engine {
	r := gin.New()

	// Gin Middleware
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// Controller
	pc := new(controllers.PingController)

	// Route List
	r.GET("/", pc.Welcome)
	r.GET("/ping", pc.Ping)
	r.GET("/ping.json", pc.PingJSON)
	r.GET("/panic", func(c *gin.Context) {
		panic("tes")
	})

	// Route Group API
	api := r.Group("/api")
	{
		api.GET("/", pc.Welcome)

	}

	// User Controller
	uc := new(controllers.UserController)
	au := api.Group("/user")
	{
		au.GET("/", uc.Index)
		au.POST("/", uc.Create)
		au.GET("/:id", uc.Detail)
		au.DELETE("/:id", uc.Delete)
	}

	return r
}
