package apitype

type CommandVoiceReq struct {
	Path string `json:"path"`
}

type AIDecisionRootCommandReq struct {
	UserCommand string `json:"usercommand"`
}
