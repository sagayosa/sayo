package api

import (
	"sayo_framework/api/module"
	"sayo_framework/api/proxy"
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

	app.Post("/proxy/plugin", proxy.Plugin(svc))
	app.Any("/proxy/:role/*url", proxy.AnyProxy(svc))

	// app.Get("/proxy/desktop/window/:way/:uuid", desktop.GetWindow(svc))
	// app.Put("/proxy/desktop/window/:way/:uuid", desktop.PutWindow(svc))
	// app.Post("/proxy/core/command/voice", core.CommandVoice(svc))
	// app.Post("/proxy/ai/chat/completions", ai.Completion(svc))
	// app.Post("/proxy/voice_recognize/voice", voicerecognize.Voice(svc))
	// app.Get("/proxy/desktop/fileselector", desktop.FileSelector(svc))
	// app.Post("/proxy/desktop/hotkey", desktop.RegisterHotKey(svc))
	// app.Post("/proxy/desktop/window", desktop.NewWindow(svc))
	// app.Get("/proxy/desktop/info/cursorposition", desktop.CursorPosition(svc))
	// app.Get("/proxy/desktop/info/workarea", desktop.WorkArea(svc))
}
