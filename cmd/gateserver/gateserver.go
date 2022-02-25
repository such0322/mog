package main

import (
	"github.com/liangdas/mqant"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	"github.com/liangdas/mqant/registry"
	"github.com/liangdas/mqant/registry/etcdv3"
	"github.com/nats-io/nats.go"
	"mog/module/web"
	"net/http"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()
	rs := etcdv3.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"127.0.0.1:2379"}
	})

	nc, err := nats.Connect("nats://127.0.0.1:4222", nats.MaxReconnects(10000))
	if err != nil {
		log.Error("nats error %v", err)
		return
	}

	app := mqant.CreateApp(
		module.Debug(true),
		module.KillWaitTTL(10*time.Second),
		module.Registry(rs),
		module.Nats(nc),
	)

	_ = app.OnConfigurationLoaded(func(app module.App) {

	})

	app.OnStartup(func(app module.App) {

	})

	err = app.Run(
		web.Module(),
	)
	if err != nil {
		log.Error(err.Error())
	}

}
