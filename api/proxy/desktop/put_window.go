package desktop

import (
	servicecontext "sayo_framework/pkg/service_context"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayoinnerhttp "github.com/grteen/sayo_utils/sayo_inner_http"
	sayoiris "github.com/grteen/sayo_utils/sayo_iris"
	"github.com/kataras/iris/v12"
)

/*
PUT /proxy/desktop/window/:way/:uuid

	json: {
		argument: interface{}
	}
*/
func PutWindow(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		var argument interface{}
		if err := ctx.ReadJSON(&argument); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}
		way := ctx.Params().GetStringDefault("way", "")
		uuid := ctx.Params().GetStringDefault("uuid", "")

		modules := module.GetInstance().GetModulesByRole(constant.RoleDesktop)
		if len(modules) == 0 {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoDesktopModule), sayoerror.ErrNoDesktopModule
		}

		result, err := sayoinnerhttp.PutWindow(modules[0].GetIPInfo(), way, uuid, argument)
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(result), nil
	})
}
