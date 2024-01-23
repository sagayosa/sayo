package main

import (
	"sayo_framework/api"
	"sayo_framework/service"

	"github.com/kataras/iris/v12"
)

// sayo_framework is only responsible for managing module configuration and distributing requests
func main() {
	app := iris.New()
	app.Use(iris.Compression)

	api.RegisterRoutes(app, service.NewServiceContext())

	app.Listen(":8080")
}
