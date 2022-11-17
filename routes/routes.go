package routes

import "github.com/gin-gonic/gin"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(c *gin.Context)
}

type Routes []Route

type routerServiceProvider struct {
	prefix string
	routes Routes
}

type routerServiceProviders []routerServiceProvider

var providers routerServiceProviders = routerServiceProviders{
	routerServiceProvider{
		prefix: "/api",
		routes: apiRoutes,
	},
}

func Register() Routes {
	var routes Routes
	for _, provider := range providers {
		for _, route := range provider.routes {
			route.Pattern = provider.prefix + route.Pattern
			routes = append(routes, route)
		}
	}
	return routes
}
