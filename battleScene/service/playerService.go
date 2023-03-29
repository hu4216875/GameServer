package service

import (
	"battleScene/model"
	"sync"
)

var (
	lock      sync.RWMutex
	playerMap map[int64]*model.PlayerData
)

func init() {
	playerMap = make(map[int64]*model.PlayerData)
}

func GetPlayerData(accountId int64) *model.PlayerData {
	lock.RLock()
	defer lock.RUnlock()

	if data, ok := playerMap[accountId]; ok {
		return data
	}
	return nil
}

func AddPlayerData(data *model.PlayerData) {
	lock.Lock()
	defer lock.Unlock()

	playerMap[data.AccountId] = data
}

func RemovePlayerData(accountId int64) {
	delete(playerMap, accountId)
}

func FindPlayerData(accountId int64) *model.PlayerData {
	lock.Lock()
	defer lock.Unlock()

	if data, ok := playerMap[accountId]; ok {
		return data
	}
	return nil
}

func RemoveGsPlayer(serverAddr string) {
	lock.Lock()
	defer lock.Unlock()
	for key, data := range playerMap {
		if data.GsServerAddr == serverAddr {
			delete(playerMap, key)
		}
	}
}
