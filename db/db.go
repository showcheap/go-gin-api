package db

import (
	"fmt"
	"go-gin-api/models"

	"github.com/jinzhu/gorm"
)

var instance *gorm.DB
var err error

// Init ...
func Init() (*gorm.DB, error) {
	instance, err = gorm.Open("mysql", "root:root@/golangapi?charset=utf8&parseTime=True&loc=Local")
	// instance.LogMode(true)

	return instance, err
}

// Instance ...
func Instance() *gorm.DB {
	return instance
}

// Migrate ...
func Migrate() {
	instance.AutoMigrate(&models.User{})
}

// Seed ...
func Seed() {
	// User Seed
	var uc int
	instance.Model(&models.User{}).Where(&models.User{Email: "admin@localhost.com"}).Count(&uc)

	if uc == 0 {
		user := models.User{Name: "Admin", Email: "admin@localhost.com"}
		user.HashPassword("admin")

		instance.Create(&user)

		fmt.Println("User created ", user)
	}
}
