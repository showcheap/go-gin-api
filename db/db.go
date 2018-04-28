package db

import (
	"github.com/jinzhu/gorm"
)

var instance *gorm.DB
var err error

// Init ...
func Init() (*gorm.DB, error) {
	instance, err = gorm.Open("mysql", "root:root@/golangapi?charset=utf8&parseTime=True&loc=Local")

	return instance, err
}

// Instance ...
func Instance() *gorm.DB {
	return instance
}
