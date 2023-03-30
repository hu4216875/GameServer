package service

import (
	"fmt"
	"github.com/name5566/leaf/log"
	"gsSelect/base"
	"gsSelect/common"
	"gsSelect/conf"
	"gsSelect/dao"
	"gsSelect/model"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	gsAddrMapLock sync.RWMutex
	gsAddrMap     map[uint32]string

	gsMapLock sync.RWMutex
	gsMap     map[uint32]*model.ServerInfo

	gsActiveMapLock sync.RWMutex
	gsActiveMap     map[uint32]*model.ServerInfo
)

func initGsService() {
	gsAddrMap = make(map[uint32]string)
	gsMap = make(map[uint32]*model.ServerInfo)
	gsActiveMap = make(map[uint32]*model.ServerInfo)

	gsList := dao.GsDao.LoadAllGs()
	for i := 0; i < len(gsList); i++ {
		gsMap[gsList[i].ServerId] = gsList[i]
	}

	discover := base.NewServiceDiscover(conf.Server.EtcdServerAddr)
	discover.ListFunc = listGsServer
	discover.UpdateFunc = updateGsServer
	discover.RemoveFunc = removeGsServer
	if err := discover.DiscoverService(common.GS_SERVER_PREFIX); err != nil {
		log.Error("InitGsService err:%v", err)
	}

	// 刷新服务器信息
	ticker := time.NewTicker(10 * time.Second)
	go refreshServerList(ticker)
}

func refreshServerList(ticker *time.Ticker) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("clientHeartCheck err:%v", err)
		}
		go refreshServerList(ticker)
	}()

	select {
	case <-ticker.C:
		gsList := dao.GsDao.LoadAllGs()
		gsAddrMapLock.Lock()
		for i := 0; i < len(gsList); i++ {
			gsMap[gsList[i].ServerId] = gsList[i]
			addActiveGs(gsList[i].ServerId)
		}
		gsAddrMapLock.Unlock()
	}
}

func listGsServer(key, value string) {
	if strings.Contains(key, common.GS_SERVER_PREFIX) {
		if serverId := addGsServer(key, value); serverId > 0 {
			addActiveGs(serverId)
		}
	}
	fmt.Println("list ", key, " value:", value)
}

func addGsServer(key, value string) uint32 {
	strAddr := getAddr(key)
	gsAddrMapLock.Lock()
	defer gsAddrMapLock.Unlock()

	serverId, err := strconv.Atoi(value)
	if err != nil {
		log.Error("listGsServer err:%v", err)
	} else {
		gsAddrMap[uint32(serverId)] = strAddr
	}

	return uint32(serverId)
}

func addActiveGs(serverId uint32) {
	serverInfo := getServerInfo(serverId)
	if serverInfo == nil {
		log.Error("addActiveGs serverId:%v not exist", serverId)
	} else {
		gsActiveMapLock.Lock()
		defer gsActiveMapLock.Unlock()
		gsActiveMap[serverId] = serverInfo
	}
}

func removeActiveGs(serverId uint32) {
	gsActiveMapLock.Lock()
	defer gsActiveMapLock.Unlock()
	delete(gsActiveMap, serverId)
}

func getServerInfo(serverId uint32) *model.ServerInfo {
	gsMapLock.Lock()
	defer gsMapLock.Unlock()

	if data, ok := gsMap[serverId]; ok {
		return data
	}
	return nil
}

func removeGs(key string) uint32 {
	strAddr := getAddr(key)

	gsAddrMapLock.Lock()
	defer gsAddrMapLock.Unlock()

	for serverId, data := range gsAddrMap {
		if data == strAddr {
			delete(gsAddrMap, serverId)
			return serverId
		}
	}
	return 0
}

func updateGsServer(key, value string) {
	if strings.Contains(key, common.GS_SERVER_PREFIX) {
		if serverId := addGsServer(key, value); serverId > 0 {
			addActiveGs(serverId)
		}
	}
	fmt.Println("update ", key, " value:", value)
}

func getAddr(key string) string {
	pos := strings.LastIndex(key, "/")
	if pos >= 0 {
		strAddr := key[pos+1:]
		strAddr = strings.Trim(strAddr, " ")
		return strAddr
	}
	return ""
}

func removeGsServer(key string) {
	if strings.Contains(key, common.GS_SERVER_PREFIX) {
		if serverId := removeGs(key); serverId > 0 {
			removeActiveGs(serverId)
		}
	}
	fmt.Println("remove ", key)
}

func getGsAddr(serverId uint32) string {
	gsAddrMapLock.RLock()
	defer gsAddrMapLock.RUnlock()

	if data, ok := gsAddrMap[serverId]; ok {
		return data
	}
	return ""
}
func selectOneRegist() string {
	gsActiveMapLock.Lock()
	defer gsActiveMapLock.Unlock()

	for serverId, data := range gsActiveMap {
		if data.RegistNum < data.LimitNum {
			data.RegistNum += 1
			return getGsAddr(serverId)
		}
	}
	return ""
}
