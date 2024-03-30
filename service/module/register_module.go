package module

import (
	"os/exec"
	"sayo_framework/pkg/constant"
	servicecontext "sayo_framework/pkg/service_context"
	servicetype "sayo_framework/pkg/type/service_type"
	"strconv"
	"time"

	"github.com/sagayosa/sayo_utils/module"
	sayoerror "github.com/sagayosa/sayo_utils/sayo_error"
	sayolog "github.com/sagayosa/sayo_utils/sayo_log"
	"github.com/sagayosa/sayo_utils/sayo_rpc/sdk"
	"github.com/sagayosa/sayo_utils/sayo_rpc/sdk/proxy"
	utils "github.com/sagayosa/sayo_utils/utils"
)

func (s *ModuleServer) RegisterModules(req *servicetype.RegisterModulesReq) (*servicetype.RegisterModulesResp, error) {
	resp := &servicetype.RegisterModulesResp{}
	for _, m := range req.Modules {
		mod, err := s.registerModule(m)
		if err != nil {
			resp.Modules = append(resp.Modules, mod)
		}
	}
	if len(resp.Modules) != 0 {
		return resp, sayoerror.ErrRegisterFailed
	}
	return resp, nil
}

func (s *ModuleServer) registerModule(m *servicetype.RegisterModuleReqModule) (*servicetype.RegisterModulesRespModule, error) {
	registerPath := utils.StringPlus(m.ModuleConfigPath, "/", constant.ModuleRegisterFile)
	config := &module.ModuleConfig{}
	if err := utils.JSON(registerPath, config); err != nil {
		return &servicetype.RegisterModulesRespModule{
			ConfigPath: m.ModuleConfigPath,
			Error:      err.Error(),
		}, err
	}

	// if config.Role == module.RolePlugin {
	// 	return s.registerPlugin(m, config)
	// }

	// port, err := utils.GetAvailablePort()
	// if err != nil {
	// 	return &servicetype.RegisterModulesRespModule{
	// 		Identifier: config.Identifier,
	// 		ConfigPath: m.ModuleConfigPath,
	// 		Error:      err.Error(),
	// 	}, err
	// }

	// mod := &module.Module{
	// 	ModuleInfo: module.ModuleInfo{
	// 		ModuleConfig: *config,
	// 		ConfigPath:   registerPath,
	// 		Address:      "127.0.0.1",
	// 		Port:         port,
	// 	},
	// }

	// if err := s.svc.ModuleCenter.RegisterModule(mod); err != nil {
	// 	return &servicetype.RegisterModulesRespModule{
	// 		Identifier: config.Identifier,
	// 		ConfigPath: m.ModuleConfigPath,
	// 		Error:      err.Error(),
	// 	}, err
	// }

	// if err := s.svc.PluginList.RegisterModule(m.ModuleConfigPath, s.svc.Cfg.PluginsList); err != nil {
	// 	return &servicetype.RegisterModulesRespModule{
	// 		Identifier: config.Identifier,
	// 		ConfigPath: m.ModuleConfigPath,
	// 		Error:      err.Error(),
	// 	}, err
	// }

	// go startModule(s.svc, config.Identifier, m.ModuleConfigPath, port)

	// return nil, nil

	return s.registerPlugin(m, config)
}

func (s *ModuleServer) registerPlugin(m *servicetype.RegisterModuleReqModule, config *module.ModuleConfig) (*servicetype.RegisterModulesRespModule, error) {
	registerPath := utils.StringPlus(m.ModuleConfigPath, "/", constant.ModuleRegisterFile)
	pluginConfig := &module.PluginConfig{}
	if err := utils.JSON(registerPath, pluginConfig); err != nil {
		return &servicetype.RegisterModulesRespModule{
			ConfigPath: m.ModuleConfigPath,
			Error:      err.Error(),
		}, err
	}

	port, err := utils.GetAvailablePort()
	if err != nil {
		return &servicetype.RegisterModulesRespModule{
			Identifier: config.Identifier,
			ConfigPath: m.ModuleConfigPath,
			Error:      err.Error(),
		}, err
	}
	plugin := &module.Module{
		ModuleInfo: module.ModuleInfo{
			ModuleConfig: *config,
			ConfigPath:   registerPath,
			Address:      "127.0.0.1",
			Port:         port,
		},
		PluginConfig: *pluginConfig,
	}

	if err := s.svc.ModuleCenter.RegisterModule(plugin); err != nil {
		return &servicetype.RegisterModulesRespModule{
			ConfigPath: m.ModuleConfigPath,
			Error:      err.Error(),
		}, err
	}

	if err := s.svc.PluginList.RegisterModule(m.ModuleConfigPath, s.svc.Cfg.PluginsList); err != nil {
		return &servicetype.RegisterModulesRespModule{
			Identifier: config.Identifier,
			ConfigPath: m.ModuleConfigPath,
			Error:      err.Error(),
		}, err
	}

	go startModule(s.svc, config.Identifier, m.ModuleConfigPath, port)
	go registerHotKey(s.svc, config)

	return nil, nil
}

func startModule(svc *servicecontext.ServiceContext, identifier string, modulePath string, port int) {
	cmd := exec.Command("cmd", "/C", ".\\process\\start_module\\module.exe", modulePath, strconv.Itoa(port), svc.GetAddr())
	if err := cmd.Run(); err != nil {
		sayolog.Err(err).Msg("identifier = %v", identifier).Info()
	}
	sayolog.Err(sayoerror.ErrModuleRestart).Msg("identifier = %v", identifier)
	time.Sleep(5 * time.Second)

	port, err := utils.GetAvailablePort()
	if err != nil {
		sayolog.Err(err).Msg("identifier = %v", identifier).Info()
	}

	modules := svc.ModuleCenter.GetModuleByIdentifier(identifier)
	if len(modules) == 0 {
		sayolog.Err(sayoerror.ErrInternalServer).Msg("restart module can not find module of identifier: %v", identifier)
	}
	module := modules[0]
	module.Port = port
	startModule(svc, identifier, modulePath, port)
}

// func startModule(svc *servicecontext.ServiceContext, modulePath string, port int) error {
// 	origin, err := os.Getwd()
// 	if err != nil {
// 		return err
// 	}
// 	if err := utils.ChangeRoutineWorkDir(modulePath); err != nil {
// 		return err
// 	}
// 	cfg := &module.ModuleConfig{}
// 	if err := utils.JSON(constant.ModuleRegisterFile, cfg); err != nil {
// 		return err
// 	}

// 	cmd := exec.Command("cmd", "/C", cfg.EntryPoint, strconv.Itoa(port), svc.GetAddr())
// 	err = cmd.Start()
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println(cmd.String())

// 	if err := utils.ChangeRoutineWorkDir(origin); err != nil {
// 		return err
// 	}
// 	return nil
// }

func registerHotKey(svc *servicecontext.ServiceContext, config *module.ModuleConfig) {
	f := func(v *module.HotKey) {
		for {
			if err := proxy.RegisterHotKey(svc.GetAddr(), &sdk.RegisterHotKeyReq{
				Identifier: config.Identifier,
				Url:        v.Url,
				Key:        v.Key,
			}); err != nil {
				sayolog.Err(sayoerror.ErrRegisterHotKeyFailed).Msg(
					"identifier = %v key = %v url = %v", config.Identifier, v.Key, v.Url).Error()
			} else {
				return
			}
			time.Sleep(5 * time.Second)
		}
	}
	for _, v := range config.HotKeys {
		go f(v)
	}
}
