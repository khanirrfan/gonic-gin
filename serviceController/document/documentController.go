package documentcontroller

import (
	"docxa/main.go/database"
	"docxa/main.go/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDocuments(c *gin.Context) {
	fmt.Println("sdjfgdsjfh")
}

func GetData(c *gin.Context) {
	fmt.Println()
	fmt.Println("Hey! I have Summoned ", c.ClientIP())
	a := &models.User{}

	db := database.GetDb()
	db.Find(a)

}
