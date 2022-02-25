package helloworld

import (
	"fmt"
	"github.com/liangdas/mqant/gate"
)

func (mo *HellWorld) say(name string) (r string, err error) {
	return fmt.Sprintf("hi %v", name), nil
}

func (mo *HellWorld) gatesay(session gate.Session,msg map[string]interface{}) (r string, err error) {
	session.Send("/gate/send/test",[]byte(fmt.Sprintf("send hi to %v", msg["name"])))
	return fmt.Sprintf("hi %v 你在网关 %v", msg["name"],session.GetServerID()), nil
}