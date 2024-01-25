package module

import (
	"context"
	baseresp "sayo_framework/pkg/base_resp"
	apitype "sayo_framework/pkg/type/api_type"
	"sayo_framework/pkg/type/cast"
	"sayo_framework/pkg/utils"
	"sayo_framework/service"
	"sayo_framework/service/module"

	"github.com/kataras/iris/v12"
)

/*
POST /module

	json: {
		modules: [
			{
				path:string 	// Prefix path for module register.json
			}
		]
	}
*/
func RegisterModule(svc *service.ServiceContext) utils.HandlerFunc {
	return utils.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		var req *apitype.RegisterModulesReq
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		resp, err := module.NewModuleServer(context.Background(), svc).RegisterModules(cast.RegisterModulesReq(req))
		if err != nil {
			return baseresp.NewBaseRespByError(err).WithData(resp), err
		}

		return baseresp.NewSuccessResp(nil), nil
	})
}
