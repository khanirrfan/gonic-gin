package main

import (
	"fmt"
	"net/http"
	"time"

	"docxa/main.go/config"
	"docxa/main.go/middleware"
	routecontroller "docxa/main.go/routeController"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	fmt.Println("Hello world!")
	r.Use(middleware.Authentication) // applied to app middleware
	routecontroller.RouteController(r)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	config.OpenDBConnection()
	server.ListenAndServe()
}
