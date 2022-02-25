package web

import (
	"context"
	"github.com/liangdas/mqant/log"
	mqrpc "github.com/liangdas/mqant/rpc"
	"io"
	"net/http"
	"time"
)

func (mo *Web) index(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	rstr, err := mqrpc.String(
		mo.Call(
			ctx,
			"helloworld",
			"/say/hi",
			mqrpc.Param(r.Form.Get("name")),
		),
	)
	log.Info("RpcCall %v , err %v", rstr, err)
	_, _ = io.WriteString(w, rstr)
}
