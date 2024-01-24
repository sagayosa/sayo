package cast

import (
	apitype "sayo_framework/pkg/type/api_type"
	servicetype "sayo_framework/pkg/type/service_type"
)

func RegisterModulesReq(req *apitype.RegisterModulesReq) *servicetype.RegisterModulesReq {
	res := &servicetype.RegisterModulesReq{}
	FillSameField(req, res)

	return res
}

func GetModulesReq(req *apitype.GetModulesReq) *servicetype.GetModulesReq {
	res := &servicetype.GetModulesReq{}
	FillSameField(req, res)

	return res
}
