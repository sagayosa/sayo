package module

import (
	"context"
	servicecontext "sayo_framework/pkg/service_context"
	apitype "sayo_framework/pkg/type/api_type"
	servicetype "sayo_framework/pkg/type/service_type"
	"sayo_framework/service/module"

	"github.com/kataras/iris/v12"
	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	"github.com/sagayosa/sayo_utils/constant"
	sayoiris "github.com/sagayosa/sayo_utils/sayo_iris"
)

/*
GET /plugin

	params: {

	}
*/
func Plugins(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.GetPluginsReq{}
		if err := ctx.ReadQuery(req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		resp, err := module.NewModuleServer(context.Background(), svc).Modules(&servicetype.GetModulesReq{
			Type: module.Role,
			Data: constant.RolePlugin,
		})
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}
		return baseresp.NewSuccessResp(resp), nil
	})
}
