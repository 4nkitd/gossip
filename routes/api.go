package routes

import v1 "fs/api/controller/v1"

var apiRoutes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		v1.Index,
	},
}
