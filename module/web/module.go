package web

import (
	"context"
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"
	"net/http"
)

var Module = func() module.Module {
	mo := new(Web)
	return mo
}

type Web struct {
	basemodule.BaseModule
}

func (mo *Web) Version() string {
	return "1.0.0"
}

func (mo *Web) GetType() string {
	return "Web"
}

func (mo *Web) OnInit(app module.App, settings *conf.ModuleSettings) {
	mo.BaseModule.OnInit(mo, app, settings)
}

func (mo *Web) startHttpServer() *http.Server {
	srv := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", mo.index)
	http.HandleFunc("/test/proto", mo.testProto)
	http.HandleFunc("/test/marshal", mo.testMarshal)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Info("Httpserver: ListenAndServe() error: %s", err)
		}
	}()
	// returning reference so caller can call Shutdown()
	return srv
}

func (mo *Web) Run(closeSig chan bool) {
	log.Info("web: starting HTTP server :8080")
	srv := mo.startHttpServer()
	<-closeSig
	log.Info("web: stopping HTTP server")
	// now close the server gracefully ("shutdown")
	// timeout could be given instead of nil as a https://golang.org/pkg/context/
	if err := srv.Shutdown(context.Background()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
	log.Info("web: done. exiting")
}

func (mo *Web) OnDestroy() {
	//一定别忘了继承
	mo.BaseModule.OnDestroy()
}
