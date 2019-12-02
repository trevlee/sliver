package rpc

import (
	"time"

	sliverpb "github.com/bishopfox/sliver/protobuf/sliver"
	"github.com/bishopfox/sliver/server/core"
	"github.com/golang/protobuf/proto"
)

func rpcRunScript(req []byte, timeout time.Duration, resp RPCResponse) {
	runScriptReq := &sliverpb.RunScriptReq{}
	err := proto.Unmarshal(req, runScriptReq)
	if err != nil {
		resp([]byte{}, err)
		return
	}
	sliver := (*core.Hive.Slivers)[runScriptReq.SliverID]
	if sliver == nil {
		resp([]byte{}, err)
		return
	}

	data, _ := proto.Marshal(&sliverpb.RunScriptReq{
		Data:     runScriptReq.Data,
		SliverID: runScriptReq.SliverID,
	})
	data, err = sliver.Request(sliverpb.MsgRunScriptReq, timeout, data)
	resp(data, err)
}
