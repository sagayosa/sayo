package module

import (
	"context"
	servicecontext "sayo_framework/pkg/service_context"
	apitype "sayo_framework/pkg/type/api_type"
	"sayo_framework/pkg/type/cast"
	"sayo_framework/service/module"

	"github.com/kataras/iris/v12"
	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	sayoiris "github.com/sagayosa/sayo_utils/sayo_iris"
)

/*
DELETE /module

	json: {
		identifiers: []
	}
*/
func UnRegisterModule(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		var req *apitype.UnRegisterModulesReq
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		module.NewModuleServer(context.Background(), svc).UnRegisterModules(cast.UnRegisterModulesReq(req))

		return baseresp.NewSuccessResp(nil), nil
	})
}
