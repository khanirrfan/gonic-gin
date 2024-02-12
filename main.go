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

	// r := gin.New()
	// fmt.Println("Hello world!")
	// r.Use(middleware.Authentication) // applied to app middleware
	// routecontroller.RouteController(r)
	// server := &http.Server{
	// 	Addr:         ":8080",
	// 	Handler:      r,
	// 	ReadTimeout:  10 * time.Second,
	// 	WriteTimeout: 10 * time.Second,
	// }
	// config.OpenDBConnection()
	// server.ListenAndServe()
}
