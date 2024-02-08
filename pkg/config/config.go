package config

type Config struct {
	PluginsList string `json:"plugins_list"`
	Port        int    `json:"port"`
}

type PluginList struct {
	Modules []*struct {
		ConfigPath string `json:"path"`
		Active     bool   `json:"active"`
	} `json:"modules"`
}
