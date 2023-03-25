package common

import (
	"errors"
	"fmt"
	"server/publicconst"
	"server/util"
	"sync"
	"time"
)

var (
	PlayerMgr = NewPlayerDataMgr()
)

type PlayerDataMgr struct {
	mutex sync.RWMutex
	data  map[int64]*PlayerData

	recycleUpdateTime uint32
}

func NewPlayerDataMgr() *PlayerDataMgr {
	ret := &PlayerDataMgr{}
	ret.data = make(map[int64]*PlayerData)
	return ret
}

// FindPlayerData 查找玩家数据
func (p *PlayerDataMgr) FindPlayerData(accountId int64) *PlayerData {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	if ret, ok := p.data[accountId]; ok {
		return ret
	}
	return nil
}

// AddPlayerData 添加玩家数据
func (p *PlayerDataMgr) AddPlayerData(playerData *PlayerData) error {
	if playerData == nil {
		return errors.New("AddPlayerData data is nil")
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.data[playerData.AccountInfo.AccountId] = playerData
	return nil
}

// DestoryPlayerData 销毁玩家数据
func (p *PlayerDataMgr) DestoryPlayerData(accountId int64) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if _, ok := p.data[accountId]; ok {
		delete(p.data, accountId)
	}
	return errors.New(fmt.Sprintf("DestoryPlayerData user:%v userdata not exitst", accountId))
}

// GetOfflinePlayer 获取所有离线的玩家
func (p *PlayerDataMgr) GetOfflinePlayer() []*PlayerData {
	curTime := uint32(time.Now().Unix())
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	var ret []*PlayerData
	maxTimeout := uint32((int)(publicconst.CLIENT_HEART_INTERVAL) * publicconst.MAX_CLIENT_HERART_NUM)
	for _, data := range p.data {
		if data.State == publicconst.Offline {
			continue
		}
		// 超时
		if curTime-data.UpdateTime > maxTimeout {
			ret = append(ret, data)
		}
	}
	return ret
}

func (p *PlayerDataMgr) RecyclePlayerData() {
	curTime := util.GetCurTime()
	if int(curTime-p.recycleUpdateTime) < publicconst.MAX_RECYCLE_PLAYER_DATA {
		return
	}
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for accountId, data := range p.data {
		if data.State == publicconst.Offline {
			if curTime-data.UpdateTime > uint32(publicconst.MAX_RECYCLE_PLAYER_DATA) {
				delete(p.data, accountId)
			}
		}
	}
	p.recycleUpdateTime = curTime
}
