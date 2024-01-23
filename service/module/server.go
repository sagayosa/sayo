package module

import (
	"context"
	"sayo_framework/service"
)

type ModuleServer struct {
	ctx context.Context
	svc *service.ServiceContext
}

func NewModuleServer(ctx context.Context, svc *service.ServiceContext) *ModuleServer {
	return &ModuleServer{
		ctx: ctx,
		svc: svc,
	}
}
