package config

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

const (
	DBUsername = "root"
	DBPassword = "root1234"
	DBHost     = "127.0.0.1"
	DBPort     = "3306"
	DBName     = "esop"
)

func OpenDBConnection() {
	var err error
	dsn := DBUsername + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println("", dsn)
	db, err = gorm.Open(sqlite.Open("esop.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting database", err)
	}
}
