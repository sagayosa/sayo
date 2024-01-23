package cast

import (
	apitype "sayo_framework/pkg/type/api_type"
	servicetype "sayo_framework/pkg/type/service_type"
)

func RegisterModuleReq(req *apitype.RegisterModuleReq) *servicetype.RegisterModuleReq {
	res := &servicetype.RegisterModuleReq{}
	FillSameField(req, res)

	return res
}

func GetModulesReq(req *apitype.GetModulesReq) *servicetype.GetModulesReq {
	res := &servicetype.GetModulesReq{}
	FillSameField(req, res)

	return res
}
