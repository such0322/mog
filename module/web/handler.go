package web

import (
	"context"
	"github.com/gogo/protobuf/proto"
	"github.com/liangdas/mqant/log"
	mqrpc "github.com/liangdas/mqant/rpc"
	rpcpb "github.com/liangdas/mqant/rpc/pb"
	"io"
	"mog/module/rpctest"
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

func (mo *Web) testProto(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	protobean := new(rpcpb.ResultInfo)
	err := mqrpc.Proto(protobean, func() (reply interface{}, errstr interface{}) {
		return mo.Call(
			ctx,
			"rpctest",     //要访问的moduleType
			"/test/proto", //访问模块中handler路径
			mqrpc.Param(&rpcpb.ResultInfo{Error: *proto.String(r.Form.Get("message"))}),
		)
	})
	log.Info("RpcCall %v , err %v", protobean, err)
	if err != nil {
		_, _ = io.WriteString(w, err.Error())
	}
	_, _ = io.WriteString(w, protobean.Error)
}

func (mo *Web) testMarshal(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*3)
	rspbean := new(rpctest.Rsp)
	err := mqrpc.Marshal(rspbean, func() (reply interface{}, errstr interface{}) {
		return mo.Call(
			ctx,
			"rpctest",       //要访问的moduleType
			"/test/marshal", //访问模块中handler路径
			mqrpc.Param(&rpctest.Req{Id: r.Form.Get("mid")}),
		)
	})
	log.Info("RpcCall %v , err %v", rspbean, err)
	if err != nil {
		_, _ = io.WriteString(w, err.Error())
	}
	_, _ = io.WriteString(w, rspbean.Msg)
}
