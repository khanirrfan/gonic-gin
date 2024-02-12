package router

import (
	"fmt"
	"net/http"
	"time"

	"docxa/main.go/config"
	"docxa/main.go/middleware"
	routecontroller "docxa/main.go/routeController"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := NewRouter()
	// router.Run(config.AppConfig.GetString("server.port"))
	server := &http.Server{
		Addr:         config.AppConfig.GetString("server.port"), // ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// config.OpenDBConnection()
	server.ListenAndServe()
}

func NewRouter() *gin.Engine {
	r := gin.New()
	fmt.Println("Hello world!")
	r.Use(middleware.Authentication) // applied to app middleware
	routecontroller.RouteController(r)
	return r
}
