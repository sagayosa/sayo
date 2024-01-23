package module

import (
	"context"
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
		path:string 	// Prefix path for module register.json
	}
*/
func RegisterModule(svc *service.ServiceContext) utils.HandlerFunc {
	return utils.IrisCtxJSONWrap(func(ctx iris.Context) (*apitype.BaseResp, error) {
		var req *apitype.RegisterModuleReq
		if err := ctx.ReadJSON(&req); err != nil {
			return apitype.NewBaseRespByError(err), err
		}

		err := module.NewModuleServer(context.Background(), svc).RegisterModule(cast.RegisterModuleReq(req))
		if err != nil {
			return apitype.NewBaseRespByError(err), err
		}

		return apitype.NewSuccessResp(), nil
	})
}
