package module

import "sayo_framework/pkg/utils"

const (
	RoleVoiceRecognize = "voice_recognize"
	RoleVoiceGenerate  = "voice_generate"
	RoleCore           = "core"
	RoleAI             = "ai"
	RoleClient         = "client"
)

type ModuleConfig struct {
	Identifier string `json:"identifier"`
	Address    string `json:"address"`
	Port       string `json:"port"`
	Role       string `json:"role"`
}

type ModuleInfo struct {
	ModuleConfig

	ConfigPath string
	SHA256     string
}

type Module struct {
	ModuleInfo
}

func (m *Module) IPInfo() string {
	return utils.StringPlus(m.Address, m.Port)
}
