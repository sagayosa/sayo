package module

const (
	RoleVoiceRecognize = "voice_recognize"
	RoleVoiceGenerate  = "voice_generate"
	RoleCore           = "core"
	RoleAI             = "ai"
	RoleClient         = "client"
)

type ModuleConfig struct {
	Identifier string
	Address    string
	Port       string
	Role       string
}

type ModuleInfo struct {
	ModuleConfig

	ConfigPath string
	SHA256     string
}

type Module struct {
	ModuleInfo
}
