package connectorcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Password string
	Username string
}

func SignIn(c *gin.Context) {
	var cred Credentials
	// Bind the JSON data to the credential struct
	if err := c.BindJSON(&cred); err != nil {
		// If there is an error in binding JSON, return a 400 Bad Request response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cred)
}

func Register(c *gin.Context) {
	// save hashed password
}

func ForgotPassword(c *gin.Context) {

}

func ChangePassword(c *gin.Context) {

}
