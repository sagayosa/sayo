package module

import (
	servicetype "sayo_framework/pkg/type/service_type"

	"github.com/grteen/sayo_utils/module"
	"github.com/grteen/sayo_utils/utils"
)

func (s *ModuleServer) Roots(req *servicetype.GetRootsReq) (resp *servicetype.GetRootsResp, err error) {
	rootMp := s.svc.ModuleCenter.GetRoots()
	resp = &servicetype.GetRootsResp{}

	for root, m := range rootMp {
		cmd := &servicetype.Command{}
		for _, declare := range m.Declare {
			if declare.Root == root {
				args := []*module.Arg{}
				for _, v := range declare.Args {
					args = append(cmd.Args, &v)
				}

				cmd.Root = root
				cmd.Args = args
				info := servicetype.ModuleInfo{}
				utils.FillSameField(m.ModuleInfo, info)
				cmd.ModuleInfo = &info
			}
		}
		resp.Cmds = append(resp.Cmds, cmd)
	}

	return resp, nil
}
