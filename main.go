package main

import "github.com/kataras/iris/v12"

// sayo_framework is only responsible for managing module configuration and distributing requests
func main() {
	app := iris.New()
	app.Use(iris.Compression)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("Hello <strong>%s</strong>!", "World")
	})

	app.Listen(":8080")
}
