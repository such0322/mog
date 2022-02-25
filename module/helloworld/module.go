package helloworld

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/log"
	"github.com/liangdas/mqant/module"
	basemodule "github.com/liangdas/mqant/module/base"
)

var Module = func() module.Module {
	mo := &HellWorld{}
	return mo
}

type HellWorld struct {
	basemodule.BaseModule
}

func (mo *HellWorld) Version() string {
	return "1.0.0"
}

func (mo *HellWorld) GetType() string {
	return "helloworld"
}

func (mo *HellWorld) OnInit(app module.App, settings *conf.ModuleSettings) {
	mo.BaseModule.OnInit(mo, app, settings)
	mo.GetServer().RegisterGO("/say/hi", mo.say)
	mo.GetServer().RegisterGO("HD_say", mo.gatesay)
	log.Info("%v模块初始化完成...", mo.GetType())
}

func (mo *HellWorld) Run(closeSig chan bool) {
	log.Info("%v模块运行中...", mo.GetType())
	log.Info("%v say hello world...", mo.GetType())
	<-closeSig
	log.Info("%v模块已停止...", mo.GetType())
}

func (mo *HellWorld) OnDestroy() {
	//一定别忘了继承
	mo.BaseModule.OnDestroy()
	log.Info("%v模块已回收...", mo.GetType())
}

func (mo *HellWorld) OnAppConfigurationLoaded(app module.App) {
	//当App初始化时调用，这个接口不管这个模块是否在这个进程运行都会调用
	mo.BaseModule.OnAppConfigurationLoaded(app)
}
