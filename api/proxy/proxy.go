package proxy

import (
	"fmt"
	"net/http"
	servicecontext "sayo_framework/pkg/service_context"

	"github.com/sagayosa/goya"
	goyawp "github.com/sagayosa/sayo_utils/goya"

	"github.com/kataras/iris/v12"
	baseresp "github.com/sagayosa/sayo_utils/base_resp"
	"github.com/sagayosa/sayo_utils/module"
	sayoerror "github.com/sagayosa/sayo_utils/sayo_error"
	sayoiris "github.com/sagayosa/sayo_utils/sayo_iris"
	frameworktypes "github.com/sagayosa/sayo_utils/types/framework"
	"github.com/sagayosa/sayo_utils/utils"
)

func AnyProxy(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
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

		method := ctx.Method()
		var resp baseresp.BaseResp
		if method == http.MethodGet {
			resp = goya.Get[baseresp.BaseResp](utils.GenerateURL(modules[0].GetIPInfo(), url), argument)
		} else if method == http.MethodPost {
			resp = goya.Post[baseresp.BaseResp](utils.GenerateURL(modules[0].GetIPInfo(), url), argument)
		} else if method == http.MethodPut {
			resp = goya.Put[baseresp.BaseResp](utils.GenerateURL(modules[0].GetIPInfo(), url), argument)
		} else if method == http.MethodDelete {
			resp = goya.Delete[baseresp.BaseResp](utils.GenerateURL(modules[0].GetIPInfo(), url), argument)
		} else {
			return baseresp.NewBaseRespByError(sayoerror.ErrUnSupportedMethod), sayoerror.ErrUnSupportedMethod
		}
		return &resp, nil
	})
}

/*
POST /proxy/plugin

	json: {
		root string
	    argvs {} struct
	}
	return data: nil
*/
func Plugin(svc *servicecontext.ServiceContext) sayoiris.HandlerFunc {
	return sayoiris.IrisCtxJSONWrap(func(ctx iris.Context) (*baseresp.BaseResp, error) {
		req := &frameworktypes.ProxyPluginReq{}
		if err := ctx.ReadJSON(&req); err != nil {
			return baseresp.NewBaseRespByError(err), err
		}

		plugins := svc.ModuleCenter.GetModuleByRoot(req.Root)
		if len(plugins) == 0 {
			return nil, sayoerror.ErrMsg(sayoerror.ErrNoPluginOfRoot, fmt.Sprintf("root = %v", req.Root))
		}
		plugin := plugins[0]

		uri := ""
		for _, r := range plugin.Declare {
			if r.Root == req.Root {
				uri = r.URL
			}
		}
		if uri == "" {
			err := sayoerror.ErrMsg(sayoerror.ErrPostPluginNoUri, fmt.Sprintf("root = %v", req.Root))
			return baseresp.NewBaseRespByError(err), err
		}

		goyawp.Post[any](plugin.GetIPInfo(), uri, req.Argvs)

		return baseresp.NewSuccessResp(nil), nil
	})
}
