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
GET /module/roots

	query: {}
*/
func Roots(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.GetRootsReq{}
		if err := ctx.ReadQuery(req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		resp, err := module.NewModuleServer(context.Background(), svc).Roots(cast.GetRootsReq(req))
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(resp), nil
	})
}
