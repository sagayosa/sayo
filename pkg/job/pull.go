package job

import (
	servicecontext "sayo_framework/pkg/service_context"
	"time"

	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayoinnerhttp "github.com/grteen/sayo_utils/sayo_inner_http"
	sayolog "github.com/grteen/sayo_utils/sayo_log"
)

func callCoreToPullCenter(svc *servicecontext.ServiceContext) error {
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

// try callCoreToPullCenter until success
func CallCoreToPullCenter(svc *servicecontext.ServiceContext) {
	f := func() {
		for {
			time.Sleep(5 * time.Second)
			if err := callCoreToPullCenter(svc); err != nil {
				sayolog.Err(err).Error(1)
				continue
			}

			return
		}
	}

	go func() {
		f()
		sayolog.Msg("%v", "core successfully pulled the module center").Info(1)
	}()
}
