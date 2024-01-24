package module

import "sayo_framework/pkg/utils"

const (
	RoleVoiceRecognize = "voice_recognize"
	RoleVoiceGenerate  = "voice_generate"
	RoleCore           = "core"
	RoleAI             = "ai"
	RoleClient         = "client"
)

var (
	RoleOptions = []string{RoleVoiceRecognize, RoleVoiceGenerate, RoleCore, RoleAI, RoleClient}
)

type ModuleConfig struct {
	Identifier string `json:"identifier"`
	Address    string `json:"address"`
	Port       string `json:"port"`
	Role       string `json:"role"`
}

type ModuleInfo struct {
	ModuleConfig

	ConfigPath string `json:"config_path"`
	// generated by the front-end to distinguish modules
	UUID string `json:"uuid"`
	// SHA256     string `json:"sha256"`
}

type Module struct {
	ModuleInfo
}

func (m *Module) IPInfo() string {
	return utils.StringPlus(m.Address, m.Port)
}
