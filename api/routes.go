package api

import (
	"sayo_framework/api/module"
	"sayo_framework/api/plugin"
	"sayo_framework/api/proxy/ai"
	"sayo_framework/api/proxy/core"
	voicerecognize "sayo_framework/api/proxy/voice_recognize"
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
	app.Post("/proxy/ai/chat/completions", ai.AIDecisionRootCommand(svc))
	app.Post("/proxy/voice_recognize/voice", voicerecognize.Voice(svc))
}
