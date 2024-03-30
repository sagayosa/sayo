package module

import (
	"sayo_framework/pkg/constant"
	servicetype "sayo_framework/pkg/type/service_type"
	"strings"

	sayolog "github.com/sagayosa/sayo_utils/sayo_log"
	"github.com/sagayosa/sayo_utils/utils"
)

func (s *ModuleServer) UnRegisterModules(req *servicetype.UnRegisterModulesReq) {
	for _, v := range req.Identifiers {
		modules := s.svc.ModuleCenter.GetModuleByIdentifier(v)
		if len(modules) == 0 {
			continue
		}
		module := modules[0]
		s.svc.ModuleCenter.UnRegisterModuleByIdentifier(v)
		if err := s.svc.PluginList.UnRegisterModule(
			strings.Replace(module.GetConfigPath(), utils.StringPlus("/", constant.ModuleRegisterFile), "", -1),
			s.svc.Cfg.PluginsList); err != nil {
			sayolog.Err(err)
		}
	}
}
