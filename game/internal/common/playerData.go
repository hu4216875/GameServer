package common

import (
	"github.com/name5566/leaf/gate"
	"server/game/internal/model"
	"server/publicconst"
	"time"
)

// PlayerData 玩家数据
type PlayerData struct {
	UserId      string
	State       publicconst.PlayerState // 玩家状态
	UpdateTime  uint32                  // 更新时间
	PlayerAgent gate.Agent
	AccountInfo *model.Account
	Items       []*model.Item // 玩家身上道具

	SceneServAddr string // 场景服务器地址
}

// NewPlayerData 玩家数据
func NewPlayerData(userId string, agent gate.Agent) *PlayerData {
	return &PlayerData{
		UserId:      userId,
		PlayerAgent: agent,
		UpdateTime:  uint32(time.Now().Unix()),
		State:       publicconst.Logining,
	}
}
