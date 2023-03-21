package game

import (
	"server/game/internal"
	"server/game/internal/common"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func GetPlayerData(userId string) *common.PlayerData {
	return common.PlayerMgr.FindPlayerData(userId)
}
