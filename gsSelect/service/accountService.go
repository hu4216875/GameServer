package service

import (
	"github.com/name5566/leaf/log"
	"gsSelect/dao"
	"sync"
	"time"
)

var (
	accountMapLock sync.RWMutex
	accountMap     map[string]*AccountAddr
)

type AccountAddr struct {
	serverAddr string
	UpdateTime uint32
}

func initAccountService() {
	accountMap = make(map[string]*AccountAddr)
	// 刷新服务器信息
	ticker := time.NewTicker(3600 * time.Second)
	go refreshAccount(ticker)
}

func getAccountByUserId(userId string) string {
	if addr := getAccount(userId); addr != nil {
		return addr.serverAddr
	}
	addrInfo := getAccountFromDb(userId)
	if addrInfo == nil {
		return ""
	}

	addAccount(userId, addrInfo)
	return addrInfo.serverAddr
}

func getAccountFromDb(userId string) *AccountAddr {
	account := dao.GsDao.FindAccount(userId)
	if account != nil {
		gsAddr := getGsAddr(account.ServerId)
		return &AccountAddr{
			serverAddr: gsAddr,
			UpdateTime: uint32(time.Now().Unix()),
		}
	}
	return nil
}

func addAccount(userId string, account *AccountAddr) {
	accountMapLock.Lock()
	defer accountMapLock.Unlock()
	accountMap[userId] = account
}

func getAccount(userId string) *AccountAddr {
	accountMapLock.RLock()
	defer accountMapLock.RUnlock()

	if data, ok := accountMap[userId]; ok {
		return data
	}
	return nil
}

func refreshAccountMap() {
	accountMapLock.Lock()
	defer accountMapLock.Unlock()

	curTime := uint32(time.Now().Unix())
	for key, data := range accountMap {
		// 超过一小时
		if curTime-data.UpdateTime >= 3600 {
			delete(accountMap, key)
		}
	}
}

func refreshAccount(ticker *time.Ticker) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("clientHeartCheck err:%v", err)
		}
		go refreshAccount(ticker)
	}()

	select {
	case <-ticker.C:
		refreshAccountMap()
	}
}
