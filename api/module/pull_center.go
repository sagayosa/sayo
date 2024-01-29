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

func PullCenter(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.PullCenterReq{}
		if err := ctx.ReadQuery(req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		resp, err := module.NewModuleServer(context.Background(), svc).PullCenter(cast.PullCenterReq(req))
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(resp), nil
	})
}
