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
	// users := []*models.User{
	// 	{Name: "Jinzhu", Age: 18, Email: "sadsad"},
	// 	{Name: "Jackson", Age: 19, Email: "sadsad"},
	// }
	a := &models.User{}

	db := database.GetDb()
	db.Find(a)
	// fmt.Println(b.Rows())
	// db.Migrator().CreateTable(&models.User{}) // uses to create dg table- should be moved to gorm config
	// db.Migrator().AddColumn(users, "email")
	// v := db.Create(&users)
	// fmt.Println(v)
}

func SetUpRoutes(r *gin.Engine) {
	admin := r.Group("/v1") // middleware
	{
		admin.GET("/getData", getData)
	}
}
