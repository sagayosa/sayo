package job

import (
	servicecontext "sayo_framework/pkg/service_context"

	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayoinnerhttp "github.com/grteen/sayo_utils/sayo_inner_http"
)

func CallCoreToPullCenter(svc *servicecontext.ServiceContext) error {
	modules := svc.ModuleCenter.GetModulesByRole(module.RoleCore)
	if len(modules) == 0 {
		return sayoerror.Msg(sayoerror.ErrCallCoreToPullCenterFailed, "%v", "There are no core modules available")
	}
	core := modules[0]
	if err := sayoinnerhttp.CallCoreToPullCenter(core.GetIPInfo()); err != nil {
		return err
	}

	return nil
}
