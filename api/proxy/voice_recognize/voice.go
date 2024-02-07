package voicerecognize

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
			return baseresp.NewBaseRespByError(sayoerror.ErrNoAIModule), sayoerror.ErrNoAIModule
		}

		result, err := sayoinnerhttp.PostVoiceRecognizeLocalFile(modules[0].GetIPInfo(), req.Path)
		if err != nil {
			return baseresp.NewBaseRespByError(sayoerror.ErrAIChatFailed), err
		}

		return baseresp.NewSuccessResp(result), nil
	})
}
