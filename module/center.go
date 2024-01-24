package module

import (
	sayoerror "sayo_framework/pkg/sayo_error"
	"sync"
)

type ModuleInterface interface {
	GetRole() string
	GetIdentifier() string
}

var (
	moduleCenterInstance *ModuleCenterSingleton = nil
	moduleCenterOnce     sync.Once
)

type ModuleCenterSingleton struct {
	roleMp   map[string][]ModuleInterface
	roleMpMu sync.Mutex

	idMp   map[string]ModuleInterface
	idMpMu sync.Mutex
}

func (s *ModuleCenterSingleton) GetModulesByRole(role string) []ModuleInterface {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.roleMp[role]
	if !ok {
		return nil
	}
	return c
}

func (s *ModuleCenterSingleton) GetModuleByIdentifier(id string) []ModuleInterface {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()

	c, ok := s.idMp[id]
	if !ok {
		return nil
	}
	return []ModuleInterface{c}
}

func (s *ModuleCenterSingleton) RegisterModule(module ModuleInterface) error {
	if err := s.registerModuleToIdentifier(module); err != nil {
		return err
	}
	s.registerModuleToRole(module)
	return nil
}

func (s *ModuleCenterSingleton) registerModuleToRole(module ModuleInterface) {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	c, ok := s.roleMp[module.GetRole()]
	if !ok {
		s.roleMp[module.GetRole()] = []ModuleInterface{module}
		return
	}

	c = append(c, module)
	s.roleMp[module.GetRole()] = c
}

func (s *ModuleCenterSingleton) registerModuleToIdentifier(module ModuleInterface) error {
	s.idMpMu.Lock()
	defer s.idMpMu.Unlock()

	_, ok := s.idMp[module.GetIdentifier()]
	if ok {
		return sayoerror.ErrDuplicateIdentifier
	}

	s.idMp[module.GetIdentifier()] = module
	return nil
}

func (s *ModuleCenterSingleton) UnRegisterModule(module ModuleInterface) {
	s.UnRegisterModuleByRole(module)
	s.UnRegisterModuleByIdentifier(module)
}

func (s *ModuleCenterSingleton) UnRegisterModuleByRole(module ModuleInterface) {
	s.roleMpMu.Lock()
	defer s.roleMpMu.Unlock()

	for key, slice := range s.roleMp {
		for idx, m := range slice {
			if m.GetIdentifier() == module.GetIdentifier() {
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

func (s *ModuleCenterSingleton) UnRegisterModuleByIdentifier(module ModuleInterface) {
	delete(s.idMp, module.GetIdentifier())
}

func (s *ModuleCenterSingleton) ClearModule() {
	s.roleMp = make(map[string][]ModuleInterface)
	s.idMp = make(map[string]ModuleInterface)
}

func newModuleCenterSingleton() *ModuleCenterSingleton {
	return &ModuleCenterSingleton{
		roleMp: make(map[string][]ModuleInterface),
		idMp:   make(map[string]ModuleInterface),
	}
}

func GetInstance() *ModuleCenterSingleton {
	moduleCenterOnce.Do(func() {
		moduleCenterInstance = newModuleCenterSingleton()
	})
	return moduleCenterInstance
}
