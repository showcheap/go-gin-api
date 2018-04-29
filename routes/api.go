package routes

import (
	"go-gin-api/controllers"
	"go-gin-api/middlewares"

	"github.com/gin-gonic/gin"
)

var secret = "MySuperSecretPassword"

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

	// Auth Controller
	ac := new(controllers.AuthController)
	lg := api.Group("/auth")
	{
		lg.POST("/login", ac.AuthenticateUser)
	}

	// User Controller
	uc := new(controllers.UserController)
	au := api.Group("/user")
	{
		au.Use(middlewares.JWTAuth())
		au.GET("/", uc.Index)
		au.POST("/", uc.Create)
		au.GET("/:id", uc.Detail)
		au.DELETE("/:id", uc.Delete)
	}

	return r
}
