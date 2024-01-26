package module

import (
	"context"
	servicecontext "sayo_framework/pkg/service_context"
)

type ModuleServer struct {
	ctx context.Context
	svc *servicecontext.ServiceContext
}

func NewModuleServer(ctx context.Context, svc *servicecontext.ServiceContext) *ModuleServer {
	return &ModuleServer{
		ctx: ctx,
		svc: svc,
	}
}
