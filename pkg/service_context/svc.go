package servicecontext

import (
	"sayo_framework/pkg/config"
	"sayo_framework/pkg/constant"
	"strconv"

	"github.com/grteen/sayo_utils/module"

	utils "github.com/grteen/sayo_utils/utils"
)

type ServiceContext struct {
	Cfg          config.Config
	ModuleCenter *module.Center
}

func (s *ServiceContext) GetAddr() string {
	return utils.StringPlus("127.0.0.1:", strconv.Itoa(s.Cfg.Port))
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
