package main

import (
	"os"
	"os/exec"
	"strconv"

	"github.com/sagayosa/sayo_utils/module"
	"github.com/sagayosa/sayo_utils/utils"
)

func startModule(modulePath string, port int, frameworkAddr string) error {
	if err := utils.ChangeRoutineWorkDir(modulePath); err != nil {
		return err
	}
	cfg := &module.ModuleConfig{}
	if err := utils.JSON("register.json", cfg); err != nil {
		return err
	}

	cmd := exec.Command("cmd", "/C", cfg.EntryPoint, strconv.Itoa(port), frameworkAddr)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	argvs := os.Args[1:]
	if len(argvs) < 3 {
		panic("at least frameworkAddr and modulePath and port are required")
	}

	port, err := strconv.Atoi(argvs[1])
	if err != nil {
		panic(err)
	}
	if err := startModule(argvs[0], port, argvs[2]); err != nil {
		panic(err)
	}
}
