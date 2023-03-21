package internal

import (
	"github.com/name5566/leaf/gate"
	"server/game/internal/common"
	"server/publicconst"
)

func init() {
	skeleton.RegisterChanRPC("NewAgent", rpcNewAgent)
	skeleton.RegisterChanRPC("CloseAgent", rpcCloseAgent)
}

func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	_ = a
}

// rpcCloseAgent 退出
func rpcCloseAgent(args []interface{}) {
	agent := args[0].(gate.Agent)
	if data := agent.UserData(); data != nil {
		if playerData := data.(*common.PlayerData); playerData != nil {
			playerData.State = publicconst.Offline
		}
	}
}
