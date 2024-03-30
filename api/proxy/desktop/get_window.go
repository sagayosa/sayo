package desktop

import (
	servicecontext "sayo_framework/pkg/service_context"

	"github.com/kataras/iris/v12"
	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	"github.com/sagayosa/sayo_utils/constant"
	"github.com/sagayosa/sayo_utils/module"
	sayoerror "github.com/sagayosa/sayo_utils/sayo_error"
	sayoiris "github.com/sagayosa/sayo_utils/sayo_iris"
	"github.com/sagayosa/sayo_utils/sayo_rpc/sdk"
)

/*
GET /proxy/desktop/window/:way/:uuid

	query: interface{}
*/
func GetWindow(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		var argument = map[string]interface{}{}
		if err := ctx.ReadQuery(&argument); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}
		way := ctx.Params().GetStringDefault("way", "")
		uuid := ctx.Params().GetStringDefault("uuid", "")

		modules := module.GetInstance().GetModulesByRole(constant.RoleDesktop)
		if len(modules) == 0 {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoDesktopModule), sayoerror.ErrNoDesktopModule
		}

		result, err := sdk.GetWindow(modules[0].GetIPInfo(), way, uuid, argument)
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(result), nil
	})
}
