package main

import (
	"fmt"
	"go-gin-api/db"

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

	r := NewRouter()

	r.Run(":8081")
}
