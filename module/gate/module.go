package gate

import (
	"github.com/liangdas/mqant/conf"
	"github.com/liangdas/mqant/gate"
	basegate "github.com/liangdas/mqant/gate/base"
	"github.com/liangdas/mqant/module"
)

var Module = func() module.Module {
	mo := &Gate{}
	return mo
}

type Gate struct {
	basegate.Gate //继承
}

func (mo *Gate) Version() string {
	return "1.0.0"
}

func (mo *Gate) GetType() string {
	return "Gate"
}

func (mo *Gate) OnInit(app module.App, settings *conf.ModuleSettings) {
	mo.Gate.OnInit(mo, app, settings,
		gate.TCPAddr(":3563"))
}
