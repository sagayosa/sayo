package module

import (
	"sayo_framework/module"
	"sayo_framework/pkg/constant"
	sayoerror "sayo_framework/pkg/sayo_error"
	servicetype "sayo_framework/pkg/type/service_type"
	"sayo_framework/pkg/utils"
)

func (s *ModuleServer) RegisterModules(req *servicetype.RegisterModulesReq) (*servicetype.RegisterModulesResp, error) {
	resp := &servicetype.RegisterModulesResp{}
	for _, m := range req.Modules {
		registerPath := utils.StringPlus(m.ModuleConfigPath, "/", constant.ModuleRegisterFile)
		config := &module.ModuleConfig{}
		if err := utils.JSON(registerPath, config); err != nil {
			resp.Modules = append(resp.Modules, struct {
				Identifier string "json:\"identifier\""
				UUID       string "json:\"uuid\""
				Error      string "json:\"error\""
			}{
				UUID:  m.UUID,
				Error: err.Error(),
			})
			continue
		}

		mod := &module.Module{
			ModuleInfo: module.ModuleInfo{
				ModuleConfig: *config,
				ConfigPath:   registerPath,
				UUID:         m.UUID,
				// SHA256:       sha,
			},
		}
		if err := module.GetInstance().RegisterModule(mod); err != nil {
			resp.Modules = append(resp.Modules, struct {
				Identifier string "json:\"identifier\""
				UUID       string "json:\"uuid\""
				Error      string "json:\"error\""
			}{
				Identifier: mod.Identifier,
				UUID:       m.UUID,
				Error:      err.Error(),
			})
			continue
		}
	}

	if len(resp.Modules) != 0 {
		return resp, sayoerror.ErrRegisterFailed
	}
	return resp, nil
}
