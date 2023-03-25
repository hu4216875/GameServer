package service

import (
	"github.com/name5566/leaf/log"
	"server/game/internal/common"
	"server/game/internal/dao"
	"server/game/internal/model"
	"server/grpc"
	"server/msg"
	"server/publicconst"
	"server/template"
	"server/util"
)

type OreService struct {
	IService
}

func NewOreService() *OreService {
	return &OreService{}
}

func (o *OreService) OnInit() {
}

func (o *OreService) OnDestory() {

}

// StartOre 开始挖矿
func (o *OreService) StartOre(accountId int64) msg.ErrCode {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	oreInfo := playerData.AccountInfo.OreInfo
	if oreInfo != nil && oreInfo.StartTime > 0 {
		return msg.ErrCode_HAS_START_ORE
	}

	oreId := template.GetSystemItemTemplate().GetOreId()
	speed := template.GetSystemItemTemplate().GetOreSpeed()
	result, err := grpc.ChanRPC.Call1("StartOre", oreId, accountId, speed)
	if err != nil {
		log.Error("AccountId:%v StartOre err:%v ", accountId, err)
		return msg.ErrCode_SYSTEM_ERROR
	}

	if result == msg.ErrCode_SUCC {
		oreInfo.Speed = speed
		oreInfo.StartTime = util.GetCurTime()
		oreInfo.OreId = oreId
		dao.OreInfoDao.UpdateOreInfo(accountId, oreInfo)
	}
	return result.(msg.ErrCode)
}

// EndOre 结束挖矿
func (o *OreService) EndOre(accountId int64) msg.ErrCode {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	oreInfo := playerData.AccountInfo.OreInfo
	if oreInfo.StartTime == 0 {
		return msg.ErrCode_NO_START_ORE
	}

	o.SettleOre(accountId, oreInfo)
	return msg.ErrCode_SUCC
}

// UpgradeOreSpeed 升级挖矿速度
func (o *OreService) UpgradeOreSpeed(accountId int64) msg.ErrCode {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	oreInfo := playerData.AccountInfo.OreInfo
	if oreInfo == nil || oreInfo.StartTime == 0 {
		return msg.ErrCode_NO_START_ORE
	}

	costItems := template.GetSystemItemTemplate().GetOreUpgradeSpeedCostItem()
	for i := 0; i < len(costItems); i++ {
		if !ServMgr.GetItemService().EnoughItem(accountId, costItems[i].ItemId, costItems[i].ItemNum) {
			return msg.ErrCode_NO_ENOUGH_ITEM
		}
	}

	oreId := template.GetSystemItemTemplate().GetOreId()
	addSpeed := template.GetSystemItemTemplate().GetOreAddSpeed()
	result, err := grpc.ChanRPC.CallN("UpgradeOreSpeed", accountId, oreId, oreInfo.Speed+addSpeed)
	if err != nil {
		log.Error("AccountId:%v UpgradeOreSpeed err:%v ", accountId, err)
		return msg.ErrCode_SYSTEM_ERROR
	}

	if result == nil {
		return msg.ErrCode_SYSTEM_ERROR
	}

	retCode := result[0].(msg.ErrCode)
	addNum := result[1].(uint32)

	if retCode == msg.ErrCode_SUCC {
		oreInfo.Speed += addSpeed
		oreInfo.StartTime = util.GetCurTime()
		dao.OreInfoDao.UpdateOreInfo(accountId, oreInfo)

		itemId := template.GetSystemItemTemplate().GetOreItemId()
		ServMgr.GetItemService().AddItem(accountId, itemId, addNum, publicconst.OreAddItem)
		for i := 0; i < len(costItems); i++ {
			ServMgr.GetItemService().CostItem(accountId, costItems[i].ItemId, costItems[i].ItemNum, publicconst.OreUpgradeSpeed)
		}
	}
	return retCode

}

// ChangeOreSpeed 改变挖矿速度
func (o *OreService) ChangeOreSpeed(accountId int64) msg.ErrCode {
	return msg.ErrCode_SUCC
}

// GetOreTotal 获取矿洞总量
func (o *OreService) GetOreTotal(accountId int64) (msg.ErrCode, uint32) {
	oreId := template.GetSystemItemTemplate().GetOreId()
	result, err := grpc.ChanRPC.Call1("GetOreTotal", oreId)
	if err != nil {
		log.Error("AccountId:%v GetOreTotal %v", accountId, err)
		return msg.ErrCode_SYSTEM_ERROR, 0
	}
	return msg.ErrCode_SUCC, result.(uint32)
}

// SettleOre 结算挖矿
func (o *OreService) SettleOre(accountId int64, oreInfo *model.OreInfo) msg.ErrCode {
	oreInfo.StartTime = 0
	oreInfo.Speed = 0
	dao.OreInfoDao.UpdateOreInfo(accountId, oreInfo)

	oreId := template.GetSystemItemTemplate().GetOreId()
	result, err := grpc.ChanRPC.Call1("EndOre", oreId, accountId)
	if err != nil {
		log.Error("AccountId:%v SettleOre %v", accountId, err)
		return msg.ErrCode_SYSTEM_ERROR
	}

	num := result.(uint32)
	if num > 0 {
		itemId := template.GetSystemItemTemplate().GetOreItemId()
		ServMgr.GetItemService().AddItem(accountId, itemId, num, publicconst.OreAddItem)
	}
	return msg.ErrCode_SUCC
}
