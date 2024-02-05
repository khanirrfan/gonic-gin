package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
	fmt.Println("Hey! I have Summoned ", c.ClientIP())
}

func SetUpRoutes(r *gin.Engine) {
	admin := r.Group("/v1") // middleware
	{
		admin.GET("/getData", getData)
	}
}
