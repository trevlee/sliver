package command

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bishopfox/sliver/client/spin"
	clientpb "github.com/bishopfox/sliver/protobuf/client"
	sliverpb "github.com/bishopfox/sliver/protobuf/sliver"
	"github.com/desertbit/grumble"
	"github.com/golang/protobuf/proto"
)

func runLuaScript(ctx *grumble.Context, rpc RPCServer) {
	if ActiveSliver.Sliver == nil {
		fmt.Printf(Warn + "Please select an active sliver via `use`\n")
		return
	}
	if len(ctx.Args) != 1 {
		fmt.Printf(Warn + "Please provide a file to load\n")
		return
	}
	scriptPath := ctx.Args[0]
	cmdTimeout := time.Duration(ctx.Flags.Int("timeout")) * time.Second
	data, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		fmt.Printf(Warn+"%v", err)
		return
	}
	scriptReqData, _ := proto.Marshal(&sliverpb.RunScriptReq{
		Data:     string(data),
		SliverID: ActiveSliver.Sliver.ID,
	})
	ctrl := make(chan bool)
	msg := "Running lua script on remote sliver ..."
	go spin.Until(msg, ctrl)
	resp := <-rpc(&sliverpb.Envelope{
		Data: scriptReqData,
		Type: clientpb.MsgRunScriptReq,
	}, cmdTimeout)
	ctrl <- true
	<-ctrl
	execResp := &sliverpb.RunScript{}
	proto.Unmarshal(resp.Data, execResp)
	if execResp.Error != "" {
		fmt.Printf(Warn+"%s", execResp.Error)
		return
	}
	if len(execResp.Result) > 0 {
		fmt.Printf("\n"+Info+"Script output:\n%s", execResp.Result)
	}
}
