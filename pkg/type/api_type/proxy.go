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

type RegisterHotKeyReq struct {
	Identifier string `json:"identifier"`
	Url        string `json:"url"`
	Key        string `json:"key"`
}

type NewWindowReq struct {
	Theme  string      `json:"theme"`
	Url    string      `json:"url"`
	Frame  bool        `json:"frame"`
	Option interface{} `json:"option"`
}

type WindowHideReq struct {
	UUID string `json:"uuid"`
}

type WindowShowReq struct {
	UUID string `json:"uuid"`
}

type WindowSetPositionReq struct {
	UUID string `json:"uuid"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type CursorPositionReq struct{}
