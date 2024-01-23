package module

import "sync"

var (
	moduleCenterInstance *ModuleCenterSingleton = nil
	moduleCenterOnce     sync.Once
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
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	for key, slice := range s.roleMp {
		for idx, m := range slice {
			if m.SHA256 == module.SHA256 || m.Identifier == module.Identifier {
				if len(slice) == 1 {
					delete(s.roleMp, key)
					return
				}

				newSlice := append(slice[:idx], slice[idx+1:]...)
				s.roleMp[key] = newSlice
			}
		}
	}
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
	moduleCenterOnce.Do(func() {
		moduleCenterInstance = newModuleCenterSingleton()
	})
	return moduleCenterInstance
}
