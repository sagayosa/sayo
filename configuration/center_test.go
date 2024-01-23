package configuration

import (
	"sayo_framework/pkg/utils"
	"testing"
)

func TestConfigCenter(t *testing.T) {
	c := GetInstance()

	var (
		r1 = ModuleConfig{Role: RoleVoiceRecognize, Address: "127.0.0.1", Port: "9877"}
		r2 = ModuleConfig{Role: RoleVoiceGenerate, Address: "192.168.131.2", Port: "8080"}
		r3 = ModuleConfig{Role: RoleAI, Address: "0.0.0.0", Port: "6379"}
		r4 = ModuleConfig{Role: RoleClient, Address: "7.7.7.7", Port: "443"}
		r5 = ModuleConfig{Role: RoleCore, Address: "256.256.256.256", Port: "80"}
		r6 = ModuleConfig{Role: RoleAI, Address: "127.0.0.1", Port: "2024"}
		r7 = ModuleConfig{Role: RoleVoiceGenerate, Address: "6.6.6.6", Port: "4048"}
		r8 = ModuleConfig{Role: RoleVoiceGenerate, Address: "5.5.5.5", Port: "9877"}
	)
	data := []struct {
		input  []*ModuleConfig
		output map[string][]*ModuleConfig
	}{
		{
			input: []*ModuleConfig{&r1, &r2, &r3, &r4, &r5},
			output: map[string][]*ModuleConfig{
				RoleVoiceRecognize: {&r1},
				RoleVoiceGenerate:  {&r2},
				RoleAI:             {&r3},
				RoleClient:         {&r4},
				RoleCore:           {&r5},
			},
		},
		{
			input: []*ModuleConfig{&r1, &r2, &r3, &r4, &r5, &r6, &r7, &r8},
			output: map[string][]*ModuleConfig{
				RoleVoiceRecognize: {&r1},
				RoleVoiceGenerate:  {&r2, &r7, &r8},
				RoleAI:             {&r3, &r6},
				RoleClient:         {&r4},
				RoleCore:           {&r5},
			},
		},
		{
			input: []*ModuleConfig{&r2, &r3, &r5, &r7},
			output: map[string][]*ModuleConfig{
				RoleVoiceGenerate: {&r2, &r7},
				RoleAI:            {&r3},
				RoleCore:          {&r5},
			},
		},
	}

	for _, d := range data {
		c.ClearModuleConfig()

		for _, in := range d.input {
			c.RegisterModuleConfig(in)
		}

		for key, val := range c.roleMp {
			ans := d.output[key]
			if !utils.CompareSlice(val, ans) {
				t.Error(utils.ComparisonFailure(ans, val))
			}
		}
	}
}
