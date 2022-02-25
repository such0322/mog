package rpctest

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	rpcpb "github.com/liangdas/mqant/rpc/pb"
)

func (mo *rpctest) testProto(req *rpcpb.ResultInfo) (*rpcpb.ResultInfo, error) {
	r := &rpcpb.ResultInfo{Error: *proto.String(fmt.Sprintf("你说: %v", req.Error))}
	return r, nil
}

func (mo *rpctest) testMarshal(req Req) (*Rsp, error) {
	r := &Rsp{Msg: fmt.Sprintf("你的ID：%v", req.Id)}
	return r, nil
}
