package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getRouteData(c *gin.Context) {
	fmt.Println("get route data")
}

func SetupRouteGetData(r *gin.Engine) {
	admin := r.Group("/v2") // apply middleware authenticate
	{
		admin.GET("/getRouteData", getRouteData)
	}
}
