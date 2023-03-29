package service

import (
	"server/game/internal/common"
	"server/game/internal/dao"
	"server/grpc"
	"server/publicconst"
	"server/util"
)

type AccountService struct {
	IService
}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (a *AccountService) OnInit() {
}

func (a *AccountService) OnDestory() {

}

func (a *AccountService) OnHeart(accountId int64) {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if playerData == nil {
		return
	}

	curTime := util.GetCurTime()
	if oreInfo := playerData.AccountInfo.OreInfo; oreInfo != nil && oreInfo.StartTime > 0 && curTime > grpc.GetOreEndTime() {
		ServMgr.GetOreService().SettleOre(accountId, oreInfo)
	}
}

func (a *AccountService) OnClose(playerData *common.PlayerData) {
	playerData.State = publicconst.Offline
	playerData.PlayerAgent.Destroy()

	accountId := playerData.AccountInfo.AccountId
	if oreInfo := playerData.AccountInfo.OreInfo; oreInfo != nil && oreInfo.StartTime > 0 {
		ServMgr.GetOreService().SettleOre(accountId, oreInfo)
	}

	if len(playerData.SceneServAddr) > 0 {
		grpc.ChanRPC.Call1("requestLeaveBattle", accountId)
		playerData.SceneServAddr = ""
	}
	dao.AccountDao.UpdateAccountLogout(accountId)
}
