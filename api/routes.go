package api

import (
	"sayo_framework/api/module"
	"sayo_framework/api/plugin"
	"sayo_framework/api/proxy/core"
	servicecontext "sayo_framework/pkg/service_context"

	"github.com/kataras/iris/v12"
)

func RegisterRoutes(app *iris.Application, svc *servicecontext.ServiceContext) {
	app.Post("/module", module.RegisterModule(svc))
	app.Get("/module", module.Modules(svc))
	app.Get("/module/role", module.ModulesByRole(svc))
	app.Get("/module/pull", module.PullCenter(svc))
	app.Get("/module/identifier", module.ModuleByIdentifier(svc))

	app.Get("/plugin", plugin.Plugins(svc))

	app.Post("/proxy/core/command/voice", core.CommandVoice(svc))
}
