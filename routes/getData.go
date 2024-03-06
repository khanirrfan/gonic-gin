package routes

import (
	"fmt"

	"docxa/main.go/database"
	"docxa/main.go/models"

	"github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
	fmt.Println()
	fmt.Println("Hey! I have Summoned ", c.ClientIP())
	a := &models.User{}

	db := database.GetDb()
	db.Find(a)

}

func SetUpRoutes(r *gin.Engine) {
	admin := r.Group("/v1") // middleware
	{
		admin.GET("/getData", getData)
	}
}
