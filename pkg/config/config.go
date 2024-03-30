package config

import (
	"sync"

	"github.com/sagayosa/sayo_utils/utils"
)

type Config struct {
	PluginsList string `json:"plugins_list"`
	Port        int    `json:"port"`
}

type PluginList struct {
	Modules []*struct {
		ConfigPath string `json:"path"`
		Active     bool   `json:"active"`
	} `json:"modules"`
	ModulesMu sync.Mutex `json:"-"`
}

func (p *PluginList) RegisterModule(ModulePath string, ConfigPath string) error {
	p.ModulesMu.Lock()
	defer p.ModulesMu.Unlock()
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

func (p *PluginList) UnRegisterModule(ModulePath string, ConfigPath string) error {
	p.ModulesMu.Lock()
	defer p.ModulesMu.Unlock()
	for _, v := range p.Modules {
		if v.ConfigPath == ModulePath {
			v.Active = false
			break
		}
	}

	return utils.JSONPersistence(ConfigPath, p)
}
