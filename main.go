package main

import (
	"fs/api/routes"
	"fs/api/system"
)

func main() {
	app := system.BootStrap()

	mapedRoutes := routes.Register()

	system.AcknowledgeRoutes(app, mapedRoutes)

	app.Run()
}
