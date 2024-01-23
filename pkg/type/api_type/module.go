package apitype

type RegisterModuleReq struct {
	ModuleConfigPath string `json:"path"`
}

type RegisterModuleResp struct {
	BaseResp
}
