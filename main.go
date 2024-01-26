package main

import (
	"sayo_framework/api"
	"sayo_framework/pkg/job"
	"sayo_framework/service"
	"sync"
	"time"

	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayolog "github.com/grteen/sayo_utils/sayo_log"

	"github.com/kataras/iris/v12"
)

var (
	svc *service.ServiceContext
)

func init() {
	svc = service.NewServiceContext()
}

func postInit(wg *sync.WaitGroup) {
	wg.Wait()
	time.Sleep(1 * time.Second)

	resp, err := job.RegisterModulesByList(svc.Cfg.ActivePluginsList, "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	if resp != nil {
		sayolog.Err(sayoerror.ErrRegisterFailed).Msg("%v", resp)
	}
}

// sayo_framework is only responsible for managing module configuration and distributing requests
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go postInit(&wg)

	app := iris.New()
	app.Use(iris.Compression)

	api.RegisterRoutes(app, svc)

	wg.Done()
	app.Listen(":8080")
}
