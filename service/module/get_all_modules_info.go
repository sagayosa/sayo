package module

import servicetype "sayo_framework/pkg/type/service_type"

func (s *ModuleServer) AllModulesInfo(req *servicetype.GetAllModulesInfoReq) (resp *servicetype.GetAllModulesInfoResp) {
	return &servicetype.GetAllModulesInfoResp{
		PluginList: s.svc.PluginList,
	}
}
