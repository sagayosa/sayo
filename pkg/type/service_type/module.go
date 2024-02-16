package servicetype

import (
	"github.com/grteen/sayo_utils/module"
)

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

type PullCenterReq struct{}

type PullCenterResp struct {
	Center *module.Center `json:"center"`
}

type GetAllModulesInfoReq struct{}

type ModuleInfo struct {
	Identifier  string `json:"identifier"`
	Active      bool   `json:"active"`
	ConfigPath  string `json:"path"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Preview     string `json:"preview"`
	Address     string `json:"address"`
}
type GetAllModulesInfoResp struct {
	ModulesInfo []*ModuleInfo `json:"modulesInfo"`
}

type UnRegisterModulesReq struct {
	Identifiers []string `json:"identifiers"`
}
