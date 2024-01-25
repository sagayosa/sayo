package service

import (
	"sayo_framework/pkg/config"
	"sayo_framework/pkg/constant"
	"sayo_framework/pkg/utils"
)

type ServiceContext struct {
	Cfg config.Config
}

func NewServiceContext() *ServiceContext {
	cfg := &config.Config{}
	if err := utils.JSON(constant.ConfigFile, cfg); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Cfg: *cfg,
	}
}
