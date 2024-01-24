package servicetype

import "sayo_framework/module"

type RegisterModulesReq struct {
	Modules []*struct {
		ModuleConfigPath string `json:"path"`
		UUID             string `json:"uuid"`
	} `json:"modules"`
}

type RegisterModulesResp struct {
	Modules []struct {
		Identifier string `json:"identifier"`
		UUID       string `json:"uuid"`
		Error      string `json:"error"`
	} `json:"errors"`
}

type GetModulesReq struct {
	Type string
	Data string
}

type GetModulesResp struct {
	Modules []module.ModuleInterface `json:"modules"`
}
