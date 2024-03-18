package desktop

import (
	servicecontext "sayo_framework/pkg/service_context"
	apitype "sayo_framework/pkg/type/api_type"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/constant"
	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayoiris "github.com/grteen/sayo_utils/sayo_iris"
	"github.com/grteen/sayo_utils/sayo_rpc/sdk"
	"github.com/grteen/sayo_utils/utils"
	"github.com/kataras/iris/v12"
)

/*
POST /proxy/desktop/window

	json: {
		theme: string,
		url: string,
		frame: boolean,

		the detail of option is in https://www.electronjs.org/zh/docs/latest/api/structures/browser-window-options
		option: interface{}
	}
*/
func NewWindow(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.NewWindowReq{}
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		modules := module.GetInstance().GetModulesByRole(constant.RoleDesktop)
		if len(modules) == 0 {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoDesktopModule), sayoerror.ErrNoDesktopModule
		}

		r := &sdk.NewWindowReq{}
		utils.FillSameField(req, r)
		uuid, err := sdk.NewWindow(modules[0].GetIPInfo(), r)
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(uuid), nil
	})
}
