package handle

import (
	"server/game/internal/common"
	"server/game/internal/service"
	"server/msg"
)

// RequestGMCommandHandle 处理GM命令
func RequestGMCommandHandle(args interface{}, playerData *common.PlayerData) {
	request := args.(*msg.RequestGMCommand)
	err := service.ServMgr.GetGmService().ProcessCommand(request.CommandId, request.Content, playerData)
	retMsg := &msg.ResponseGMCommand{
		Result: int32(err),
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}
