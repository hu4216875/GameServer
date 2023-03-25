package handle

import (
	"server/game/internal/common"
	"server/game/internal/service"
	"server/msg"
)

// RequestLoadItemHandle 请求加载道具
func RequestLoadItemHandle(args interface{}, playerData *common.PlayerData) {
	err, items := service.ServMgr.GetItemService().LoadItems(playerData.AccountInfo.AccountId)

	retMsg := &msg.ResponseLoadItem{
		Result: int32(err),
	}
	if len(items) > 0 {
		retMsg.Items = append(retMsg.Items, service.ToProtocolItems(items)...)
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}
