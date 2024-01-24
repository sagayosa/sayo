package servicetype

import "sayo_framework/module"

type RegisterModuleReqModule struct {
	ModuleConfigPath string `json:"path"`
}

type RegisterModulesReq struct {
	Modules []*RegisterModuleReqModule `json:"modules"`
}

type RegisterModulesRespModule struct {
	Identifier string `json:"identifier"`
	ConfigPath string `json:"path"`
	Error      string `json:"error"`
}
type RegisterModulesResp struct {
	Modules []*RegisterModulesRespModule `json:"errors"`
}

type GetModulesReq struct {
	Type string
	Data string
}

type GetModulesResp struct {
	Modules []module.ModuleInterface `json:"modules"`
}
