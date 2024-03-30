package voicerecognize

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
POST /proxy/voice_recognize/voice

	json: {
		path string
	}
*/
func Voice(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.VoiceReq{}
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		modules := module.GetInstance().GetModulesByRole(constant.RoleVoiceRecognize)
		if len(modules) == 0 {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoVoiceRecognizeModule), sayoerror.ErrNoVoiceRecognizeModule
		}

		result, err := sdk.PostVoiceRecognizeLocalFile(modules[0].GetIPInfo(), req.Path)
		if err != nil {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoVoiceRecognizeModule), err
		}

		return baseresp.NewSuccessResp(result), nil
	})
}
