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
	PluginList   *config.PluginList
	// PluginsCmd   []*exec.Cmd
}

func (s *ServiceContext) GetAddr() string {
	return utils.StringPlus("127.0.0.1:", strconv.Itoa(s.Cfg.Port))
}

// func (s *ServiceContext) RegisterCmd(c *exec.Cmd) {
// 	s.PluginsCmd = append(s.PluginsCmd, c)
// 	go c.Wait()
// }

func NewServiceContext() *ServiceContext {
	cfg := &config.Config{}
	if err := utils.JSON(constant.ConfigFile, cfg); err != nil {
		panic(err)
	}
	list := &config.PluginList{}
	if err := utils.JSON(cfg.PluginsList, list); err != nil {
		panic(err)
	}

	return &ServiceContext{
		Cfg:          *cfg,
		ModuleCenter: module.GetInstance(),
		PluginList:   list,
		// PluginsCmd:   make([]*exec.Cmd, 0),
	}
}
