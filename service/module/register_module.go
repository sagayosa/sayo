package module

import (
	"sayo_framework/pkg/constant"
	servicetype "sayo_framework/pkg/type/service_type"

	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	utils "github.com/grteen/sayo_utils/utils"
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

	if config.Role == module.RolePlugin {
		return s.registerPlugin(m, config)
	}

	port, err := utils.GetAvailablePort()
	if err != nil {
		return &servicetype.RegisterModulesRespModule{
			Identifier: config.Identifier,
			ConfigPath: m.ModuleConfigPath,
			Error:      err.Error(),
		}, err
	}

	mod := &module.Module{
		ModuleInfo: module.ModuleInfo{
			ModuleConfig: *config,
			ConfigPath:   registerPath,
			Address:      "127.0.0.1",
			Port:         port,
		},
	}

	if err := s.svc.ModuleCenter.RegisterModule(mod); err != nil {
		return &servicetype.RegisterModulesRespModule{
			Identifier: config.Identifier,
			ConfigPath: m.ModuleConfigPath,
			Error:      err.Error(),
		}, err
	}

	return nil, nil
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
	plugin := &module.Plugin{
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

	return nil, nil
}
