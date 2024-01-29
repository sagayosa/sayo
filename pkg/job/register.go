package job

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"sayo_framework/pkg/constant"
	apitype "sayo_framework/pkg/type/api_type"
	servicetype "sayo_framework/pkg/type/service_type"

	baseresp "github.com/grteen/sayo_utils/base_resp"
	"github.com/grteen/sayo_utils/module"
	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayolog "github.com/grteen/sayo_utils/sayo_log"
	utils "github.com/grteen/sayo_utils/utils"
)

type ActivePlugin struct {
	ModulePaths []string `json:"modules"`
}

func RegisterModulesByList(listPath string, address string, center *module.Center) (*servicetype.RegisterModulesResp, error) {
	active := &ActivePlugin{}
	if err := utils.JSON(listPath, active); err != nil {
		return nil, err
	}

	resp := &servicetype.RegisterModulesResp{Modules: make([]*servicetype.RegisterModulesRespModule, 0)}
	for _, p := range active.ModulePaths {
		res, err := sendRequest(p, address)
		if err != nil {
			sayolog.Err(err)
			continue
		}
		startModules(p, center)

		if res != nil {
			resp.Modules = append(resp.Modules, res.Modules...)
		}
	}

	return resp, nil
}

func startModules(active string, center *module.Center) {
	start := func(p string) {
		err := func() error {
			if err := utils.ChangeRoutineWorkDir(p); err != nil {
				return err
			}
			cfg := &module.ModuleConfig{}
			if err := utils.JSON(constant.ModuleRegisterFile, cfg); err != nil {
				return err
			}

			mods := center.GetModuleByIdentifier(cfg.Identifier)
			if len(mods) == 0 {
				return fmt.Errorf("no such identifier: %v", cfg.Identifier)
			}
			mod := mods[0]

			info := mod.GetIPInfo()
			_, port, err := utils.SplitIPInfo(info)
			if err != nil {
				return err
			}

			cmd := exec.Command("cmd", "/C", cfg.EntryPoint, port, "127.0.0.1")
			_, err = cmd.Output()
			if err != nil {
				return err
			}

			return nil
		}()
		if err != nil {
			sayolog.Err(sayoerror.ErrRunModulesFailed).Msg(err.Error())
		}
	}

	go start(active)
}

func sendRequest(active string, address string) (res *servicetype.RegisterModulesResp, err error) {
	req := &apitype.RegisterModulesReq{}

	req.Modules = append(req.Modules, &servicetype.RegisterModuleReqModule{
		ModuleConfigPath: active,
	})

	code, body, err := utils.Post(utils.StringPlus("http://", address, "/module"), req)
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
