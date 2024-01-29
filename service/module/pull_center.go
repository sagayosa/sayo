package module

import (
	servicetype "sayo_framework/pkg/type/service_type"
)

func (s *ModuleServer) PullCenter(req *servicetype.PullCenterReq) (resp *servicetype.PullCenterResp, err error) {
	resp = &servicetype.PullCenterResp{}
	resp.Center = s.svc.ModuleCenter

	return resp, nil
}
