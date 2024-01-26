package plugin

import (
	"context"
	mod "sayo_framework/module"
	apitype "sayo_framework/pkg/type/api_type"
	servicetype "sayo_framework/pkg/type/service_type"
	"sayo_framework/service/module"

	servicecontext "sayo_framework/pkg/service_context"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	sayoiris "github.com/grteen/sayo_utils/sayo_iris"

	"github.com/kataras/iris/v12"
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
			Data: mod.RolePlugin,
		})
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}
		return baseresp.NewSuccessResp(resp), nil
	})
}
