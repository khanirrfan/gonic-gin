package routes

import (
	documentcontroller "docxa/main.go/serviceController/document"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(r *gin.Engine) {
	admin := r.Group("/v1") // middleware
	{
		admin.GET("/getData", documentcontroller.GetDocuments)
	}
}
