package module

import (
	"context"
	apitype "sayo_framework/pkg/type/api_type"
	"sayo_framework/pkg/type/cast"
	servicetype "sayo_framework/pkg/type/service_type"
	"sayo_framework/service"
	"sayo_framework/service/module"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	utils "github.com/grteen/sayo_utils/utils"

	"github.com/kataras/iris/v12"
)

/*
GET /module

	params: {
		type:string 	// The type of request, with optional types including [role, identifier]
		data:string 	// The data according to the type
	}
*/
func Modules(svc *service.ServiceContext) utils.HandlerFunc {
	return utils.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.GetModulesReq{}
		if err := ctx.ReadQuery(req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		resp, err := module.NewModuleServer(context.Background(), svc).Modules(cast.GetModulesReq(req))
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(resp), nil
	})
}

/*
GET /module/role

	params: {
		role:string
	}
*/
func ModulesByRole(svc *service.ServiceContext) utils.HandlerFunc {
	return utils.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.GetModulesByRoleReq{}
		if err := ctx.ReadQuery(req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		resp, err := module.NewModuleServer(context.Background(), svc).Modules(&servicetype.GetModulesReq{
			Type: module.Role,
			Data: req.Role,
		})
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(resp), nil
	})
}

/*
GET /module/identifier

	params: {
		identifier:string
	}
*/
func ModuleByIdentifier(svc *service.ServiceContext) utils.HandlerFunc {
	return utils.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.GetModuleByIdentifierReq{}
		if err := ctx.ReadQuery(req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		resp, err := module.NewModuleServer(context.Background(), svc).Modules(&servicetype.GetModulesReq{
			Type: module.Identifier,
			Data: req.Identifier,
		})
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(resp), nil
	})
}
