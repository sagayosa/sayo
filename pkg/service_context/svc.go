package servicecontext

import (
	"sayo_framework/pkg/config"
	"sayo_framework/pkg/constant"

	"github.com/grteen/sayo_utils/module"

	utils "github.com/grteen/sayo_utils/utils"
)

type ServiceContext struct {
	Cfg          config.Config
	ModuleCenter *module.Center
}

func NewServiceContext() *ServiceContext {
	cfg := &config.Config{}
	if err := utils.JSON(constant.ConfigFile, cfg); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Cfg:          *cfg,
		ModuleCenter: module.GetInstance(),
	}
}
