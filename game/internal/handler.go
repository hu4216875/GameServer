package internal

import (
	"github.com/name5566/leaf/gate"
	"reflect"
	"server/game/internal/common"
	"server/game/internal/handle"
	"server/msg"
)

type HandleFunc func(args interface{}, playerData *common.PlayerData)

func init() {
	skeleton.RegisterChanRPC("Regist", handle.RpcRegist)
	skeleton.RegisterChanRPC("Login", handle.RpcLogin)

	handler(&msg.RequestLoadItem{}, makeHandleMsg(handle.RequestLoadItemHandle))
	handler(&msg.RequestGMCommand{}, makeHandleMsg(handle.RequestGMCommandHandle))
	handler(&msg.RequestClientHeart{}, makeHandleMsg(handle.RequestClientHeartHandle))
	handler(&msg.RequestLogout{}, makeHandleMsg(handle.RequestLogoutHandle))

	handler(&msg.RequestOreInfo{}, makeHandleMsg(handle.RequestOreInfoHandle))
	handler(&msg.RequestOreTotal{}, makeHandleMsg(handle.RequestOreTotalHandle))
	handler(&msg.RequestStartOre{}, makeHandleMsg(handle.RequestStartOreHandle))
	handler(&msg.RequestEndOre{}, makeHandleMsg(handle.RequestEndOreHandle))
	handler(&msg.RequestUpgradeOreSpeed{}, makeHandleMsg(handle.RequestUpgradeOreSpeedHandle))
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func makeHandleMsg(f HandleFunc) func(args []interface{}) {
	return func(args []interface{}) {
		agent := args[1].(gate.Agent)
		userData := agent.UserData()
		if userData == nil {
			kickMsg := &msg.ResponseKickOut{Result: int32(msg.ErrCode_USER_NOT_LOGIN)}
			agent.WriteMsg(kickMsg)
			agent.Destroy()
			return
		}
		playerData := userData.(*common.PlayerData)
		f(args[0], playerData)
	}
}
