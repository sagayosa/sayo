package servicetype

import "sayo_framework/module"

type RegisterModuleReq struct {
	ModuleConfigPath string
}

type GetModulesReq struct {
	Type string
	Data string
}

type GetModulesResp struct {
	Modules []*module.Module `json:"modules"`
}
