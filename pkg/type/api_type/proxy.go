package apitype

type CommandVoiceReq struct {
	Path string `json:"path"`
}

type AIDecisionRootCommandReq struct {
	UserCommand string `json:"usercommand"`
}

type CompletionReq struct {
	Content string `json:"content"`
}

type VoiceReq struct {
	Path string `json:"path"`
}

type PluginReq struct {
	Root  string      `json:"root"`
	Argvs interface{} `json:"argvs"`
}

type FileSelectorReq struct{}
