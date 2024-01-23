package configuration

import "sync"

var (
	configCenterInstance *ConfigCenterSingleton = nil
	configCenterOnce     sync.Once
)

type ConfigCenterSingleton struct {
	roleMp   map[string][]*ModuleConfig
	roleMpMu sync.Mutex
}

func (s *ConfigCenterSingleton) GetModuleConfigByRole(role string) []*ModuleConfig {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.roleMp[role]
	if !ok {
		return nil
	}
	return c
}

func (s *ConfigCenterSingleton) RegisterModuleConfig(moduleConfig *ModuleConfig) {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.roleMp[moduleConfig.Role]
	if !ok {
		s.roleMp[moduleConfig.Role] = []*ModuleConfig{moduleConfig}
		return
	}

	c = append(c, moduleConfig)
	s.roleMp[moduleConfig.Role] = c
}

func (s *ConfigCenterSingleton) ClearModuleConfig() {
	s.roleMp = make(map[string][]*ModuleConfig)
}

func newConfigCenterSingleton() *ConfigCenterSingleton {
	return &ConfigCenterSingleton{
		roleMp: make(map[string][]*ModuleConfig),
	}
}

func GetInstance() *ConfigCenterSingleton {
	configCenterOnce.Do(func() {
		configCenterInstance = newConfigCenterSingleton()
	})
	return configCenterInstance
}
