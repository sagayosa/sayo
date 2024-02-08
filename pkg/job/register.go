package job

import (
	"encoding/json"
	"fmt"
	"net/http"
	servicecontext "sayo_framework/pkg/service_context"
	apitype "sayo_framework/pkg/type/api_type"
	servicetype "sayo_framework/pkg/type/service_type"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayolog "github.com/grteen/sayo_utils/sayo_log"
	utils "github.com/grteen/sayo_utils/utils"
)

func RegisterModulesByList(svc *servicecontext.ServiceContext) (*servicetype.RegisterModulesResp, error) {
	resp := &servicetype.RegisterModulesResp{Modules: make([]*servicetype.RegisterModulesRespModule, 0)}
	for _, p := range svc.PluginList.Modules {
		if !p.Active {
			continue
		}
		res, err := sendRequest(svc, p.ConfigPath)
		if err != nil {
			sayolog.Err(err)
			continue
		}

		if res != nil {
			resp.Modules = append(resp.Modules, res.Modules...)
		}
	}

	return resp, nil
}

// func startModules(svc *servicecontext.ServiceContext, active string) {
// 	start := func(p string) {
// 		err := func() error {
// 			if err := utils.ChangeRoutineWorkDir(p); err != nil {
// 				return err
// 			}
// 			cfg := &module.ModuleConfig{}
// 			if err := utils.JSON(constant.ModuleRegisterFile, cfg); err != nil {
// 				return err
// 			}

// 			mods := svc.ModuleCenter.GetModuleByIdentifier(cfg.Identifier)
// 			if len(mods) == 0 {
// 				return fmt.Errorf("no such identifier: %v", cfg.Identifier)
// 			}
// 			mod := mods[0]

// 			info := mod.GetIPInfo()
// 			_, port, err := utils.SplitIPInfo(info)
// 			if err != nil {
// 				return err
// 			}

// 			cmd := exec.Command("cmd", "/C", cfg.EntryPoint, port, svc.GetAddr())
// 			err = cmd.Start()
// 			if err != nil {
// 				return err
// 			}
// 			fmt.Println(cmd.String())

// 			return nil
// 		}()
// 		if err != nil {
// 			sayolog.Err(sayoerror.ErrRunModulesFailed).Msg(err.Error())
// 		}
// 	}

// 	start(active)
// }

func sendRequest(svc *servicecontext.ServiceContext, active string) (res *servicetype.RegisterModulesResp, err error) {
	req := &apitype.RegisterModulesReq{}

	req.Modules = append(req.Modules, &servicetype.RegisterModuleReqModule{
		ModuleConfigPath: active,
	})

	code, body, err := utils.Post(utils.StringPlus("http://", svc.GetAddr(), "/module"), req)
	if err != nil {
		return
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("statusCode = %v", code)
	}

	resp := &baseresp.BaseResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return
	}

	if resp.Code != sayoerror.SuccessCode {
		body := &servicetype.RegisterModulesResp{}
		err = utils.UnMarshalUnknownAny(resp.Data, body)
		if err != nil {
			return
		}

		return body, nil
	}

	return nil, nil
}
