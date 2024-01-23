package module

import (
	"sayo_framework/module"
	servicetype "sayo_framework/pkg/type/service_type"
)

const (
	Role       = "role"
	Identifier = "identifier"
)

func (s *ModuleServer) Modules(req *servicetype.GetModulesReq) (resp *servicetype.GetModulesResp, err error) {
	resp = &servicetype.GetModulesResp{}
	if req.Type == Role {
		modules := module.GetInstance().GetModulesByRole(req.Data)
		resp.Modules = modules
		return
	}
	if req.Type == Identifier {
		modules := module.GetInstance().GetModuleByIdentifier(req.Data)
		resp.Modules = modules
		return
	}

	return
}
