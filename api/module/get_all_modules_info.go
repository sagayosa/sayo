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
GET /module/info
*/
func AllModulesInfo(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.GetAllModulesInfoReq{}
		if err := ctx.ReadQuery(req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		resp, err := module.NewModuleServer(context.Background(), svc).AllModulesInfo(cast.GetAllModulesInfoReq(req))
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(resp.ModulesInfo), nil
	})
}
