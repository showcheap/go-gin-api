package controllers

import (
	"fmt"
	"go-gin-api/db"
	"go-gin-api/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthController ...
type AuthController struct{}

var secret = "MySuperSecretPassword"

type userRequest struct {
}

// AuthenticateUser ...
func (a AuthController) AuthenticateUser(c *gin.Context) {

	// Post Param
	var form struct {
		Email    string `form:"email" json:"email"`
		Password string `form:"password" json:"password"`
	}

	c.Bind(&form)

	// Query On DB
	db := db.Instance()

	var user models.User
	if err := db.Model(&models.User{}).Where(&models.User{Email: form.Email}).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(404, gin.H{"message": "User not found!"})
		return
	}

	if err := user.CheckPassword(form.Password); err != nil {
		c.AbortWithStatusJSON(401, gin.H{"message": "Email & Password doesn't match"})
		return
	}

	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"Id":   user.ID,
		"name": user.Name,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		fmt.Println("Unable generate token:", err)
		c.AbortWithStatusJSON(500, gin.H{"message": "Could not generate token!"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString, "exp": time.Now().Add(time.Hour * 1)})
}
