package module

import (
	servicetype "sayo_framework/pkg/type/service_type"

	"github.com/sagayosa/sayo_utils/module"
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
					arg := v
					args = append(args, &arg)
				}
				cmd.Root = root
				cmd.Args = args
				cmd.ModuleInfo = &servicetype.ModuleInfo{
					Identifier:  m.Identifier,
					Active:      true,
					ConfigPath:  m.ConfigPath,
					Name:        m.Name,
					Description: m.Description,
					Author:      m.Author,
					Preview:     m.Preview,
					Address:     m.Address,
				}
			}
		}
		resp.Cmds = append(resp.Cmds, cmd)
	}

	return resp, nil
}
