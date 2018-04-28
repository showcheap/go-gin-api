package controllers

import "github.com/gin-gonic/gin"
import jwt "github.com/dgrijalva/jwt-go"
import "time"

type AuthController struct{}

var secret = "MySuperSecretPassword"

// AuthenticateUser ...
func (a AuthController) AuthenticateUser(c *gin.Context) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"Id":  "Sucipto",
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	}

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(500, gin.H{"message": "Could not generate token!"})
	}

	c.JSON(200, gin.H{"token": tokenString, "exp": time.Now().Add(time.Hour * 1)})
}
