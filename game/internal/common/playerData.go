package common

import (
	"github.com/name5566/leaf/gate"
	"server/publicconst"
	"time"
)

// PlayerData 玩家数据
type PlayerData struct {
	UserId      string
	AccountId   int64
	State       publicconst.PlayerState // 玩家状态
	UpdateTime  uint32                  // 更新时间
	PlayerAgent gate.Agent
}

// NewPlayerData 玩家数据
func NewPlayerData(userId string, accountId int64, agent gate.Agent) *PlayerData {
	return &PlayerData{
		UserId:      userId,
		AccountId:   accountId,
		PlayerAgent: agent,
		UpdateTime:  uint32(time.Now().Unix()),
		State:       publicconst.Logining,
	}
}
