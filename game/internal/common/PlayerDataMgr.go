package common

import (
	"errors"
	"fmt"
	"server/publicconst"
	"sync"
	"time"
)

var (
	PlayerMgr = NewPlayerDataMgr()
)

type PlayerDataMgr struct {
	mutex sync.RWMutex
	data  map[string]*PlayerData
}

func NewPlayerDataMgr() *PlayerDataMgr {
	ret := &PlayerDataMgr{}
	ret.data = make(map[string]*PlayerData)
	return ret
}

// FindPlayerData 查找玩家数据
func (p *PlayerDataMgr) FindPlayerData(userId string) *PlayerData {
	p.mutex.RLock()
	defer p.mutex.RUnlock()
	if ret, ok := p.data[userId]; ok {
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
	p.data[playerData.UserId] = playerData
	return nil
}

// DestoryPlayerData 销毁玩家数据
func (p *PlayerDataMgr) DestoryPlayerData(userId string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if _, ok := p.data[userId]; ok {
		delete(p.data, userId)
	}
	return errors.New(fmt.Sprintf("DestoryPlayerData user:%v userdata not exitst", userId))
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
