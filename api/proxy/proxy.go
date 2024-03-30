package proxy

import (
	servicecontext "sayo_framework/pkg/service_context"

	"github.com/kataras/iris/v12"
	"github.com/sagayosa/goya"
	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	"github.com/sagayosa/sayo_utils/module"
	sayoerror "github.com/sagayosa/sayo_utils/sayo_error"
	sayoiris "github.com/sagayosa/sayo_utils/sayo_iris"
	"github.com/sagayosa/sayo_utils/utils"
)

func GetProxy(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		var argument any
		if err := ctx.ReadBody(&argument); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}
		role := ctx.Params().GetStringDefault("role", "")
		url := ctx.Params().GetStringDefault("url", "")

		modules := module.GetInstance().GetModulesByRole(role)
		if len(modules) == 0 {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoModule), sayoerror.ErrNoModule
		}

		return baseresp.NewSuccessResp(goya.Get[any](utils.StringPlus("http://", modules[0].GetIPInfo(), "/", url), nil)), nil
	})
}
