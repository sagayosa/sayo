package module

import (
	"sayo_framework/pkg/constant"
	servicetype "sayo_framework/pkg/type/service_type"

	"github.com/grteen/sayo_utils/module"
	utils "github.com/grteen/sayo_utils/utils"
)

func (s *ModuleServer) AllModulesInfo(req *servicetype.GetAllModulesInfoReq) (resp *servicetype.GetAllModulesInfoResp, err error) {
	res := []*servicetype.ModuleInfo{}
	for _, v := range s.svc.PluginList.Modules {
		registerPath := utils.StringPlus(v.ConfigPath, "/", constant.ModuleRegisterFile)
		config := &module.ModuleConfig{}
		if err = utils.JSON(registerPath, config); err != nil {
			return
		}

		info := &servicetype.ModuleInfo{}
		utils.FillSameField(config, info)
		info.Active = v.Active
		info.ConfigPath = v.ConfigPath

		res = append(res, info)
	}

	return &servicetype.GetAllModulesInfoResp{ModulesInfo: res}, nil
}
