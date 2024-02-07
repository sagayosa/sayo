package ai

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
POST /proxy/ai/chat/completions

	json: {
		content string
	}
*/
func Completion(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &apitype.CompletionReq{}
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		modules := module.GetInstance().GetModulesByRole(constant.RoleAI)
		if len(modules) == 0 {
			return baseresp.NewBaseRespByError(sayoerror.ErrNoAIModule), sayoerror.ErrNoAIModule
		}

		result, err := sayoinnerhttp.PostAICompletion(modules[0].GetIPInfo(), req.Content)
		if err != nil {
			return baseresp.NewBaseRespByError(sayoerror.ErrAIChatFailed), err
		}

		return baseresp.NewSuccessResp(result), nil
	})
}