package module

import (
	"sayo_framework/module"
	"sayo_framework/pkg/constant"
	servicetype "sayo_framework/pkg/type/service_type"
	"sayo_framework/pkg/utils"
)

func (s *ModuleServer) RegisterModule(req *servicetype.RegisterModuleReq) error {
	registerPath := utils.StringPlus(req.ModuleConfigPath, "/", constant.ModuleRegisterFile)

	config := &module.ModuleConfig{}
	if err := utils.JSON(registerPath, config); err != nil {
		return err
	}

	// sha, err := utils.SHA256(registerPath)
	// if err != nil {
	// 	return err
	// }

	mod := module.Module{
		ModuleInfo: module.ModuleInfo{
			ModuleConfig: *config,
			ConfigPath:   registerPath,
			// SHA256:       sha,
		},
	}
	module.GetInstance().RegisterModule(&mod)

	return nil
}
