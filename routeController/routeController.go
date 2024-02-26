package routecontroller

import (
	"docxa/main.go/routes"

	"github.com/gin-gonic/gin"
)

func RouteController(r *gin.Engine) {

	routes.SetUpRoutes(r)
	routes.SetupRouteGetData(r)
}
