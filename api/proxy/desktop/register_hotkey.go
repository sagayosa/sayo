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
	"github.com/kataras/iris/v12"
)

/*
POST /proxy/desktop/hotkey

	json: {
		identifier: string,
		key: string,
		url: string
	}
*/
func RegisterHotKey(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.RegisterHotKeyReq{}
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		modules := module.GetInstance().GetModulesByRole(constant.RoleDesktop)
		if len(modules) == 0 {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoDesktopModule), sayoerror.ErrNoDesktopModule
		}

		err := sdk.RegisterHotKey(modules[0].GetIPInfo(), req.Identifier, &module.HotKey{Key: req.Key, Url: req.Url})
		if err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(nil), nil
	})
}
