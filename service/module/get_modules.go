package module

import (
	servicetype "sayo_framework/pkg/type/service_type"

	"github.com/grteen/sayo_utils/module"

	sayoerror "github.com/grteen/sayo_utils/sayo_error"
)

const (
	Role       = "role"
	Identifier = "identifier"
)

func (s *ModuleServer) Modules(req *servicetype.GetModulesReq) (resp *servicetype.GetModulesResp, err error) {
	if err = check(req); err != nil {
		return
	}
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

func check(req *servicetype.GetModulesReq) error {
	if req.Type == Role {
		return nil
	} else if req.Type == Identifier {
		return nil
	} else {
		return sayoerror.ErrUnknownType
	}
}
