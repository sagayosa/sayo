package module

import "sync"

var (
	configCenterInstance *ModuleCenterSingleton = nil
	configCenterOnce     sync.Once
)

type ModuleCenterSingleton struct {
	roleMp   map[string][]*Module
	roleMpMu sync.Mutex
}

func (s *ModuleCenterSingleton) GetModulesByRole(role string) []*Module {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.roleMp[role]
	if !ok {
		return nil
	}
	return c
}

func (s *ModuleCenterSingleton) RegisterModule(module *Module) {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.roleMp[module.Role]
	if !ok {
		s.roleMp[module.Role] = []*Module{module}
		return
	}

	c = append(c, module)
	s.roleMp[module.Role] = c
}

func (s *ModuleCenterSingleton) UnRegisterModule(module *Module) {

}

func (s *ModuleCenterSingleton) ClearModule() {
	s.roleMp = make(map[string][]*Module)
}

func newModuleCenterSingleton() *ModuleCenterSingleton {
	return &ModuleCenterSingleton{
		roleMp: make(map[string][]*Module),
	}
}

func GetInstance() *ModuleCenterSingleton {
	configCenterOnce.Do(func() {
		configCenterInstance = newModuleCenterSingleton()
	})
	return configCenterInstance
}
