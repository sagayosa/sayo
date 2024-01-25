package job

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"sayo_framework/module"
	baseresp "sayo_framework/pkg/base_resp"
	"sayo_framework/pkg/constant"
	sayoerror "sayo_framework/pkg/sayo_error"
	sayolog "sayo_framework/pkg/sayo_log"
	apitype "sayo_framework/pkg/type/api_type"
	servicetype "sayo_framework/pkg/type/service_type"
	"sayo_framework/pkg/utils"
	"strconv"
)

type ActivePlugin struct {
	ModulePaths []string `json:"modules"`
}

func RegisterModulesByList(listPath string, address string) (*servicetype.RegisterModulesResp, error) {
	active := &ActivePlugin{}
	if err := utils.JSON(listPath, active); err != nil {
		return nil, err
	}

	resp, err := sendRequest(active, address)
	startModules(active)

	return resp, err
}

func startModules(active *ActivePlugin) {
	for _, p := range active.ModulePaths {
		start := func(p string) {
			err := func() error {
				if err := utils.ChangeRoutineWorkDir(p); err != nil {
					return err
				}
				cfg := &module.ModuleConfig{}
				if err := utils.JSON(constant.ModuleRegisterFile, cfg); err != nil {
					return err
				}

				mods := module.GetInstance().GetModuleByIdentifier(cfg.Identifier)
				if len(mods) == 0 {
					return fmt.Errorf("no such identifier: %v", cfg.Identifier)
				}
				mod := mods[0]
				_, port := mod.GetIPInfo()

				cmd := exec.Command("cmd", "/C", cfg.EntryPoint, strconv.Itoa(port))
				_, err := cmd.Output()
				if err != nil {
					return err
				}

				return nil
			}()
			if err != nil {
				sayolog.Err(sayoerror.ErrRunModulesFailed).Msg(err.Error())
			}
		}

		go start(p)
	}
}

func sendRequest(active *ActivePlugin, address string) (res *servicetype.RegisterModulesResp, err error) {
	req := &apitype.RegisterModulesReq{}

	for _, p := range active.ModulePaths {
		req.Modules = append(req.Modules, &servicetype.RegisterModuleReqModule{
			ModuleConfigPath: p,
		})
	}

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
