package module

import (
	"context"
	apitype "sayo_framework/pkg/type/api_type"
	"sayo_framework/pkg/type/cast"
	"sayo_framework/service/module"

	servicecontext "sayo_framework/pkg/service_context"

	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	sayoiris "github.com/sagayosa/sayo_utils/sayo_iris"

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
func RegisterModule(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
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
