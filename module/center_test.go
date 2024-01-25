package module

import (
	"sayo_framework/pkg/utils"
	"testing"
)

var (
	r1 = Module{ModuleInfo: ModuleInfo{ModuleConfig: ModuleConfig{Identifier: "1", Role: RoleVoiceRecognize}}} // Address: "127.0.0.1", Port: "9877",
	r2 = Module{ModuleInfo: ModuleInfo{ModuleConfig: ModuleConfig{Identifier: "-2", Role: RoleVoiceGenerate}}} // Address: "192.168.131.2", Port: "8080"
	r3 = Module{ModuleInfo: ModuleInfo{ModuleConfig: ModuleConfig{Identifier: "3", Role: RoleAI}}}             // Address: "0.0.0.0", Port: "6379"
	r4 = Module{ModuleInfo: ModuleInfo{ModuleConfig: ModuleConfig{Identifier: "-4", Role: RoleClient}}}        // Address: "7.7.7.7", Port: "443"
	r5 = Module{ModuleInfo: ModuleInfo{ModuleConfig: ModuleConfig{Identifier: "5", Role: RoleCore}}}           // Address: "256.256.256.256", Port: "80"
	r6 = Module{ModuleInfo: ModuleInfo{ModuleConfig: ModuleConfig{Identifier: "-6", Role: RoleAI}}}            // Address: "127.0.0.1", Port: "2024"
	r7 = Module{ModuleInfo: ModuleInfo{ModuleConfig: ModuleConfig{Identifier: "7", Role: RoleVoiceGenerate}}}  // Address: "6.6.6.6", Port: "4048"
	r8 = Module{ModuleInfo: ModuleInfo{ModuleConfig: ModuleConfig{Identifier: "-8", Role: RoleVoiceGenerate}}} // Address: "5.5.5.5", Port: "9877"
)

func TestRegisterModule(t *testing.T) {
	c := GetInstance()

	data := []struct {
		input  []ModuleInterface
		output map[string][]ModuleInterface
	}{
		{
			input: []ModuleInterface{&r1, &r2, &r3, &r4, &r5},
			output: map[string][]ModuleInterface{
				RoleVoiceRecognize: {&r1},
				RoleVoiceGenerate:  {&r2},
				RoleAI:             {&r3},
				RoleClient:         {&r4},
				RoleCore:           {&r5},
			},
		},
		{
			input: []ModuleInterface{&r1, &r2, &r3, &r4, &r5, &r6, &r7, &r8},
			output: map[string][]ModuleInterface{
				RoleVoiceRecognize: {&r1},
				RoleVoiceGenerate:  {&r2, &r7, &r8},
				RoleAI:             {&r3, &r6},
				RoleClient:         {&r4},
				RoleCore:           {&r5},
			},
		},
		{
			input: []ModuleInterface{&r2, &r3, &r5, &r7},
			output: map[string][]ModuleInterface{
				RoleVoiceGenerate: {&r2, &r7},
				RoleAI:            {&r3},
				RoleCore:          {&r5},
			},
		},
	}

	for _, d := range data {
		c.ClearModule()

		for _, in := range d.input {
			if err := c.RegisterModule(in); err != nil {
				t.Error(err)
			}
		}

		for key, val := range c.roleMp {
			ans := d.output[key]
			if !utils.CompareSlice(val, ans) {
				t.Error(utils.ComparisonFailure(ans, val))
			}
		}
	}
}

func TestUnRegisterModule(t *testing.T) {
	c := GetInstance()

	data := []struct {
		input struct {
			register   []*Module
			unRegister []*Module
		}
		output map[string][]ModuleInterface
	}{
		{
			input: struct {
				register   []*Module
				unRegister []*Module
			}{
				register:   []*Module{&r1, &r2, &r3, &r4, &r5},
				unRegister: []*Module{&r1, &r2, &r3},
			},
			output: map[string][]ModuleInterface{
				RoleClient: {&r4},
				RoleCore:   {&r5},
			},
		},
		{
			input: struct {
				register   []*Module
				unRegister []*Module
			}{
				register:   []*Module{&r1, &r2, &r3, &r4, &r5, &r7, &r8},
				unRegister: []*Module{&r1, &r2, &r3, &r7},
			},
			output: map[string][]ModuleInterface{
				RoleClient:        {&r4},
				RoleCore:          {&r5},
				RoleVoiceGenerate: {&r8},
			},
		},
	}

	for _, d := range data {
		c.ClearModule()

		for _, in := range d.input.register {
			c.RegisterModule(in)
		}
		for _, in := range d.input.unRegister {
			c.UnRegisterModule(in)
		}

		for key, val := range c.roleMp {
			ans := d.output[key]
			if !utils.CompareSlice(val, ans) {
				t.Error(utils.ComparisonFailure(ans, val))
			}

			for _, v := range val {
				moduels := c.GetModuleByIdentifier(v.GetIdentifier())
				if moduels == nil {
					t.Error(utils.MissingFailure(v.GetIdentifier()))
				}
			}
		}
	}
}

func TestGetModuleByIdentifier(t *testing.T) {
	c := GetInstance()

	data := []struct {
		input  string
		output []ModuleInterface
	}{
		{
			input:  "1",
			output: []ModuleInterface{&r1},
		},
		{
			input:  "-2",
			output: []ModuleInterface{&r2},
		},
		{
			input:  "2",
			output: nil,
		},
	}

	register := []*Module{&r1, &r2, &r3, &r4, &r5}

	for _, m := range register {
		c.RegisterModule(m)
	}

	for _, d := range data {
		res := c.GetModuleByIdentifier(d.input)

		if !utils.CompareSlice(res, d.output) {
			t.Error(utils.ComparisonFailure(d.output, res))
		}
	}
}
