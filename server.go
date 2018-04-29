package main

import (
	"fmt"
	"go-gin-api/db"
	"go-gin-api/routes"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {

	// Get DB Connection
	c, err := db.Init()

	if err != nil {
		fmt.Println("Connect database error: ", err)
	}

	// Close DB on program exit
	defer c.Close()

	// AutoMigrate
	db.Migrate()
	db.Seed()

	r := routes.NewRouter()

	r.Run(":8081")
}
