package connectorcontroller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignIn(c *gin.Context) {
	fmt.Println("c")
	c.JSON(http.StatusOK, "sdfsdf")
	// use jwt,
	// validation - for multiple times wrong entering credentials
	//

}

func Register(c *gin.Context) {
	// save hashed password
}

func ForgotPassword(c *gin.Context) {

}

func ChangePassword(c *gin.Context) {

}
