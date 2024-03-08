package api

import (
	"sayo_framework/api/module"
	"sayo_framework/api/proxy/ai"
	"sayo_framework/api/proxy/core"
	"sayo_framework/api/proxy/desktop"
	"sayo_framework/api/proxy/plugin"
	voicerecognize "sayo_framework/api/proxy/voice_recognize"
	servicecontext "sayo_framework/pkg/service_context"

	"github.com/kataras/iris/v12"
)

func RegisterRoutes(app *iris.Application, svc *servicecontext.ServiceContext) {
	app.Post("/module", module.RegisterModule(svc))
	app.Delete("/module", module.UnRegisterModule(svc))
	app.Get("/module", module.Modules(svc))
	app.Get("/module/role", module.ModulesByRole(svc))
	app.Get("/module/pull", module.PullCenter(svc))
	app.Get("/module/identifier", module.ModuleByIdentifier(svc))
	app.Get("/plugin", module.Plugins(svc))
	app.Get("/module/info", module.AllModulesInfo(svc))
	app.Get("/module/roots", module.Roots(svc))

	app.Post("/proxy/core/command/voice", core.CommandVoice(svc))
	app.Post("/proxy/ai/chat/completions", ai.Completion(svc))
	app.Post("/proxy/voice_recognize/voice", voicerecognize.Voice(svc))
	app.Post("/proxy/plugin", plugin.Plugin(svc))
	app.Get("/proxy/desktop/fileselector", desktop.FileSelector(svc))
	app.Post("/proxy/desktop/hotkey", desktop.RegisterHotKey(svc))
	app.Post("/proxy/desktop/window", desktop.NewWindow(svc))
	app.Put("/proxy/desktop/window/hide", desktop.WindowHide(svc))
}
