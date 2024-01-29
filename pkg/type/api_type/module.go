package apitype

import servicetype "sayo_framework/pkg/type/service_type"

type RegisterModulesReq struct {
	Modules []*servicetype.RegisterModuleReqModule `json:"modules"`
}

// type RegisterModulesResp struct {
// 	BaseResp
// }

type GetModulesReq struct {
	Type string `param:"type"`
	Data string `param:"data"`
}

type GetModulesByRoleReq struct {
	Role string `param:"role"`
}

type GetModuleByIdentifierReq struct {
	Identifier string `param:"identifier"`
}

type PullCenterReq struct {
}

// type GetModulesResp struct {
// 	BaseResp
// }
