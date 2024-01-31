package main

import (
	"sayo_framework/api"
	"sayo_framework/pkg/job"
	servicecontext "sayo_framework/pkg/service_context"
	"strconv"
	"sync"
	"time"

	sayoerror "github.com/grteen/sayo_utils/sayo_error"
	sayolog "github.com/grteen/sayo_utils/sayo_log"
	"github.com/grteen/sayo_utils/utils"

	"github.com/kataras/iris/v12"
)

var (
	svc *servicecontext.ServiceContext
)

func init() {
	svc = servicecontext.NewServiceContext()
}

func postInit(wg *sync.WaitGroup) {
	wg.Wait()
	time.Sleep(1 * time.Second)

	resp, err := job.RegisterModulesByList(svc)
	if err != nil {
		panic(err)
	}
	if resp != nil {
		sayolog.Err(sayoerror.ErrRegisterFailed).Msg("%v", resp)
	}

	job.CallCoreToPullCenter(svc)
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
	app.Listen(utils.StringPlus(":", strconv.Itoa(svc.Cfg.Port)))
}
