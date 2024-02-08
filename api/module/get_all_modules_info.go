package module

import (
	"context"
	servicecontext "sayo_framework/pkg/service_context"
	apitype "sayo_framework/pkg/type/api_type"
	"sayo_framework/pkg/type/cast"
	"sayo_framework/service/module"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	sayoiris "github.com/grteen/sayo_utils/sayo_iris"
	"github.com/kataras/iris/v12"
)

/*
GET /module/info
*/
func AllModulesInfo(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.GetAllModulesInfoReq{}
		if err := ctx.ReadQuery(req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		resp := module.NewModuleServer(context.Background(), svc).AllModulesInfo(cast.GetAllModulesInfoReq(req))

		return baseresp.NewSuccessResp(resp), nil
	})
}
