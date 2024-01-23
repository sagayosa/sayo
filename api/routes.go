package api

import (
	"sayo_framework/api/module"
	"sayo_framework/service"

	"github.com/kataras/iris/v12"
)

func RegisterRoutes(app *iris.Application, svc *service.ServiceContext) {
	app.Post("/module", module.RegisterModule(svc))
}
