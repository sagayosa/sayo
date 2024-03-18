package plugin

import (
	"fmt"
	servicecontext "sayo_framework/pkg/service_context"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayoinnerhttp "github.com/grteen/sayo_utils/sayo_inner_http"
	sayoiris "github.com/grteen/sayo_utils/sayo_iris"
	"github.com/grteen/sayo_utils/sayo_rpc/sdk"
	"github.com/kataras/iris/v12"
)

/*
POST /proxy/plugin

	json: {
		root string
	    argvs {} struct
	}
*/
func Plugin(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &sdk.AIDecisionResp{}
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		plugins := svc.ModuleCenter.GetModuleByRoot(req.Root)
		if len(plugins) == 0 {
			return nil, sayoerror.ErrMsg(sayoerror.ErrNoPluginOfRoot, fmt.Sprintf("root = %v", req.Root))
		}
		plugin := plugins[0]

		if err := sayoinnerhttp.PostPlugin(plugin, req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		return baseresp.NewSuccessResp(nil), nil
	})
}
