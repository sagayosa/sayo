package config

import "github.com/grteen/sayo_utils/utils"

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

func (p *PluginList) RegisterModule(ModulePath string, ConfigPath string) error {
	find := false
	for _, v := range p.Modules {
		if v.ConfigPath == ModulePath {
			v.Active = true
			find = true
			break
		}
	}

	if !find {
		p.Modules = append(p.Modules, &struct {
			ConfigPath string "json:\"path\""
			Active     bool   "json:\"active\""
		}{ConfigPath: ModulePath, Active: true})
	}
	return utils.JSONPersistence(ConfigPath, p)
}

func (p *PluginList) UnRegosterModule(ModulePath string, ConfigPath string) error {
	for _, v := range p.Modules {
		if v.ConfigPath == ConfigPath {
			v.Active = false
			break
		}
	}

	return utils.JSONPersistence(ConfigPath, p)
}
