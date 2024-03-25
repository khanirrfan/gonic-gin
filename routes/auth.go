package routes

import (
	connectorcontroller "docxa/main.go/connectorController"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authRoutes := r.Group("/user")
	{
		authRoutes.POST("/login", connectorcontroller.SignIn)
		authRoutes.POST("/register", connectorcontroller.Register)
		authRoutes.POST("/forgotPassword", connectorcontroller.ForgotPassword)
		authRoutes.POST("/changePassword", connectorcontroller.ChangePassword)
	}
}
