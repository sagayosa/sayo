package module

const (
	RoleVoiceRecognize = "voice_recognize"
	RoleVoiceGenerate  = "voice_generate"
	RoleCore           = "core"
	RoleAI             = "ai"
	RoleClient         = "client"
	RolePlugin         = "plugin"
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
}

type Module struct {
	ModuleInfo
}

func (m *ModuleInfo) GetIdentifier() string {
	return m.Identifier
}

func (m *ModuleInfo) GetRole() string {
	return m.Role
}
