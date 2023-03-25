package game

import (
	"server/game/internal"
	"server/game/internal/common"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func GetPlayerData(accountId int64) *common.PlayerData {
	return common.PlayerMgr.FindPlayerData(accountId)
}
