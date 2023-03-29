package service

import (
	"github.com/name5566/leaf/log"
	"server/game/internal/common"
	"server/grpc"
	"server/msg"
	"server/publicconst"
)

type BattleService struct {
	IService
}

func NewBattleService() *BattleService {
	return &BattleService{}
}

func (b *BattleService) OnInit() {
	grpc.Module.SetSceneCloseFunc(OnSceneClose)
}

func (b *BattleService) OnDestory() {

}

func (b *BattleService) RequestEnterBattle(accountId int64) int32 {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if len(playerData.SceneServAddr) > 0 {
		return int32(msg.ErrCode_USER_IN_BATTLE)
	}
	result, err := grpc.ChanRPC.CallN("requestEnterBattle", accountId)
	if err != nil || len(result) == 0 {
		log.Error("AccountId:%v RequestEnterBattle err:%v ", accountId, err)
		return int32(msg.ErrCode_SYSTEM_ERROR)
	}

	errCode := result[0].(int32)
	if errCode == int32(msg.ErrCode_SUCC) {
		servAddr := result[1].(string)
		playerData.SceneServAddr = servAddr
	}
	return errCode
}

func (b *BattleService) RequestLeaveBattle(accountId int64) int32 {
	playerData := common.PlayerMgr.FindPlayerData(accountId)
	if len(playerData.SceneServAddr) == 0 {
		return int32(msg.ErrCode_USER_NOT_IN_BATTLE)
	}
	result, err := grpc.ChanRPC.Call1("requestLeaveBattle", accountId)
	if err != nil {
		log.Error("AccountId:%v RequestEnterBattle err:%v ", accountId, err)
		return int32(msg.ErrCode_SYSTEM_ERROR)
	}
	return result.(int32)
}

func OnSceneClose(sceneAddr string) {
	players := common.PlayerMgr.GetScenePlayer(sceneAddr)
	retMsg := &msg.ResponseLeaveBattle{Result: int32(msg.ErrCode_SUCC)}
	for i := 0; i < len(players); i++ {
		players[i].SceneServAddr = ""
		if players[i].State == publicconst.Online {
			players[i].PlayerAgent.WriteMsg(retMsg)
		}
	}
}
