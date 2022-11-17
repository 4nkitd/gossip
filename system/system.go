package system

import (
	"fs/api/routes"

	"github.com/gin-gonic/gin"
)

func BootStrap() *gin.Engine {
	return gin.Default()
}

func AcknowledgeRoutes(app *gin.Engine, routes routes.Routes) {
	for _, route := range routes {
		app.Handle(route.Method, route.Pattern, route.HandlerFunc)
	}
}
