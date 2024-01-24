package apitype

type RegisterModuleReq struct {
	ModuleConfigPath string `json:"path"`
}

type RegisterModuleResp struct {
	BaseResp
}

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

type GetModulesResp struct {
	BaseResp
}
