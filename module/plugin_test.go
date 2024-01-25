package module

import (
	"sayo_framework/pkg/utils"
	"testing"
)

var (
	p1 = &Plugin{
		ModuleInfo: ModuleInfo{
			ModuleConfig: ModuleConfig{
				Identifier: "1",
				Role:       RolePlugin,
				// Address:    "127.0.0.1",
				// Port:       "9877",
			},
		},
		PluginConfig: PluginConfig{
			Declare: []Declare{
				{
					Root:        "exampleRoot1",
					Description: "This is the description for the first PluginConfig",
					Args: []Arg{
						{
							Name:        "arg1",
							Description: "Description for arg1",
						},
					},
				},
			},
		}}

	p2 = &Plugin{
		ModuleInfo: ModuleInfo{
			ModuleConfig: ModuleConfig{
				Identifier: "2",
				Role:       RolePlugin,
				// Address:    "192.168.0.1",
				// Port:       "8765",
			},
		},
		PluginConfig: PluginConfig{
			Declare: []Declare{
				{
					Root:        "exampleRoot2",
					Description: "This is the description for the second PluginConfig",
					Args: []Arg{
						{
							Name:        "arg2",
							Description: "Description for arg2",
						},
						{
							Name:        "arg3",
							Description: "Description for arg3",
						},
					},
				},
			},
		},
	}

	p3 = &Plugin{
		ModuleInfo: ModuleInfo{
			ModuleConfig: ModuleConfig{
				Identifier: "3",
				Role:       RolePlugin,
				// Address:    "10.0.0.1",
				// Port:       "7654",
			},
		},
		PluginConfig: PluginConfig{
			Declare: []Declare{
				{
					Root:        "exampleRoot3",
					Description: "This is the description for the third PluginConfig",
					Args: []Arg{
						{
							Name:        "arg4",
							Description: "Description for arg4",
						},
						{
							Name:        "arg5",
							Description: "Description for arg5",
						},
						{
							Name:        "arg6",
							Description: "Description for arg6",
						},
					},
				},
			},
		},
	}

	p4 = &Plugin{
		ModuleInfo: ModuleInfo{
			ModuleConfig: ModuleConfig{
				Identifier: "4",
				Role:       RolePlugin,
				// Address:    "172.16.0.1",
				// Port:       "5432",
			},
		},
		PluginConfig: PluginConfig{
			Declare: []Declare{
				{
					Root:        "exampleRoot4",
					Description: "This is the description for the fourth PluginConfig",
					Args: []Arg{
						{
							Name:        "arg7",
							Description: "Description for arg7",
						},
						{
							Name:        "arg8",
							Description: "Description for arg8",
						},
						{
							Name:        "arg9",
							Description: "Description for arg9",
						},
						{
							Name:        "arg10",
							Description: "Description for arg10",
						},
					},
				},
			},
		},
	}
)

func TestPlugin(t *testing.T) {
	c := GetInstance()

	data := []struct {
		input   []ModuleInterface
		exclude []ModuleInterface
		output  []ModuleInterface
	}{
		{
			input:   []ModuleInterface{p1, p2},
			exclude: []ModuleInterface{},
			output:  []ModuleInterface{p1, p2},
		},
		{
			input:   []ModuleInterface{p1, p2, p3, p4},
			exclude: []ModuleInterface{p2, p4},
			output:  []ModuleInterface{p1, p3},
		},
	}

	for _, d := range data {
		c.ClearModule()
		for _, p := range d.input {
			if err := c.RegisterModule(p); err != nil {
				t.Error(err)
			}
		}
		for _, p := range d.exclude {
			c.UnRegisterModule(p)
		}
		res := c.GetModulesByRole(RolePlugin)
		if !utils.CompareSlice(res, d.output) {
			t.Error(utils.ComparisonFailure(d.output, res))
		}

		for _, p := range res {
			_, ok := p.(*Plugin)
			if !ok {
				t.Error(utils.ConvetFailure(p))
			}
		}
	}
}
