package core

import (
	servicecontext "sayo_framework/pkg/service_context"
	apitype "sayo_framework/pkg/type/api_type"

	"github.com/kataras/iris/v12"
	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	"github.com/sagayosa/sayo_utils/constant"
	"github.com/sagayosa/sayo_utils/module"
	sayoerror "github.com/sagayosa/sayo_utils/sayo_error"
	sayoiris "github.com/sagayosa/sayo_utils/sayo_iris"
	"github.com/sagayosa/sayo_utils/sayo_rpc/sdk"
)

/*
POST /proxy/core/command/voice

	json: {
		path string
	}
*/
func CommandVoice(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.CommandVoiceReq{}
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		modules := module.GetInstance().GetModulesByRole(constant.RoleCore)
		if len(modules) == 0 {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoCoreModule), sayoerror.ErrNoCoreModule
		}

		if err := sdk.CoreVoiceCommand(modules[0].GetIPInfo(), req.Path); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(nil), nil
	})
}
