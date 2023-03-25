package handle

import (
	"server/game/internal/common"
	"server/game/internal/service"
	"server/msg"
	"server/template"
)

// RequestOreInfoHandle 请求挖矿信息
func RequestOreInfoHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseOreInfo{
		Result: int32(msg.ErrCode_SUCC),
	}

	oreInfo := playerData.AccountInfo.OreInfo
	if oreInfo != nil {
		retMsg.OreId = oreInfo.OreId
		retMsg.StartTime = oreInfo.StartTime
		retMsg.Speed = oreInfo.Speed
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}

// RequestOreTotalHandle 请求矿洞总量
func RequestOreTotalHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseOreTotal{}
	err, total := service.ServMgr.GetOreService().GetOreTotal(playerData.AccountInfo.AccountId)
	retMsg.Result = int32(err)
	if err == msg.ErrCode_SUCC {
		retMsg.Total = total
		retMsg.OreId = template.GetSystemItemTemplate().GetOreId()
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}

// RequestStartOreHandle 开始挖矿
func RequestStartOreHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseStartOre{
		Result: int32(service.ServMgr.GetOreService().StartOre(playerData.AccountInfo.AccountId)),
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}

// RequestEndOreHandle 结束挖矿
func RequestEndOreHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseEndOre{
		Result: int32(service.ServMgr.GetOreService().EndOre(playerData.AccountInfo.AccountId)),
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}

// RequestUpgradeOreSpeedHandle 升级挖矿速度
func RequestUpgradeOreSpeedHandle(args interface{}, playerData *common.PlayerData) {
	retMsg := &msg.ResponseUpgradeOreSpeed{
		Result: int32(service.ServMgr.GetOreService().UpgradeOreSpeed(playerData.AccountInfo.AccountId)),
	}
	playerData.PlayerAgent.WriteMsg(retMsg)
}
