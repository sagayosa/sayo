package desktop

import (
	servicecontext "sayo_framework/pkg/service_context"
	apitype "sayo_framework/pkg/type/api_type"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayoinnerhttp "github.com/grteen/sayo_utils/sayo_inner_http"
	sayoiris "github.com/grteen/sayo_utils/sayo_iris"
	"github.com/kataras/iris/v12"
)

/*
PUT /proxy/desktop/window/show

	json: {
		uuid string
	}
*/
func WindowShow(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.WindowHideReq{}
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		modules := module.GetInstance().GetModulesByRole(constant.RoleDesktop)
		if len(modules) == 0 {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoDesktopModule), sayoerror.ErrNoDesktopModule
		}

		if err := sayoinnerhttp.WindowShow(modules[0].GetIPInfo(), req.UUID); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(nil), nil
	})
}
