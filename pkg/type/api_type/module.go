package apitype

type RegisterModulesReq struct {
	Modules []*struct {
		ModuleConfigPath string `json:"path"`
		UUID             string `json:"uuid"`
	} `json:"modules"`
}

type RegisterModulesResp struct {
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
