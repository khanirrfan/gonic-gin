package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getRouteData(c *gin.Context) {
	fmt.Println("safhgsdjhfgsdjfd")
}

func SetupRouteGetData(r *gin.Engine) {
	admin := r.Group("/v2") // apply middleware authenticate
	{
		admin.GET("/getRouteData", getRouteData)
	}
}
