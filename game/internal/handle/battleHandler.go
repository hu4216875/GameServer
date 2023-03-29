package handle

import (
	"server/game/internal/common"
	"server/game/internal/service"
	"server/msg"
)

// RequestEnterBattleHandle 进入战斗
func RequestEnterBattleHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseEnterBattle{}
	retMsg.Result = service.ServMgr.GetBattleService().RequestEnterBattle(playerData.AccountInfo.AccountId)
	playerData.PlayerAgent.WriteMsg(retMsg)
}

// RequestLeaveBattleHandle 离开战斗
func RequestLeaveBattleHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseLeaveBattle{}
	retMsg.Result = int32(service.ServMgr.GetBattleService().RequestLeaveBattle(playerData.AccountInfo.AccountId))
	playerData.PlayerAgent.WriteMsg(retMsg)
}
