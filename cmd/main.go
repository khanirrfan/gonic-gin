package main

import (
	"docxa/main.go/config"
	"docxa/main.go/database"
	"docxa/main.go/router"
)

func main() {
	config.Init()
	config.AppConfig = config.GetConfig()
	database.OpenDBConnection()
	router.Init()

}
