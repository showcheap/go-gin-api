package main

import (
	"fmt"
	"go-gin-api/db"
	"go-gin-api/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	db, err := db.Init()

	if err != nil {
		fmt.Println("Connect database error: ", err)
	}

	// Close DB on program exit
	defer db.Close()

	// AutoMigrate
	db.AutoMigrate(&models.User{})

	r := NewRouter()

	r.Run(":8081")
}
