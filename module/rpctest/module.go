package rpctest

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"
	"github.com/liangdas/mqant/server"
	"time"
)

var Module = func() module.Module {
	mo := &rpctest{}
	return mo
}

type rpctest struct {
	basemodule.BaseModule
}

func (mo *rpctest) Version() string {
	return "1.0.0"
}

func (mo *rpctest) GetType() string {
	return "rpctest"
}

func (mo *rpctest) OnInit(app module.App, settings *conf.ModuleSettings) {
	mo.BaseModule.OnInit(mo, app, settings,
		server.RegisterInterval(15*time.Second),
		server.RegisterTTL(30*time.Second),
		//server.ID("rpcnode001"),
	)
	//mo.GetServer().Options().Metadata["serID"] = "rpc001"

	mo.GetServer().RegisterGO("/test/proto", mo.testProto)
	mo.GetServer().RegisterGO("/test/marshal", mo.testMarshal)
}

func (mo *rpctest) Run(closeSig chan bool) {
	<-closeSig
}

func (mo *rpctest) OnDestroy() {
	mo.BaseModule.OnDestroy()
}
