package module

import (
	"context"
	servicecontext "sayo_framework/pkg/service_context"
)

type ModuleServer struct {
	ctx context.Context
	svc *servicecontext.ServiceContext
	// registerLimiter chan struct{}
}

// func (s *ModuleServer) registerAccess() {
// 	s.registerLimiter <- struct{}{}
// }

// func (s *ModuleServer) registerLeft() {
// 	<-s.registerLimiter
// }

func NewModuleServer(ctx context.Context, svc *servicecontext.ServiceContext) *ModuleServer {
	return &ModuleServer{
		ctx: ctx,
		svc: svc,
		// registerLimiter: make(chan struct{}, 1),
	}
}
