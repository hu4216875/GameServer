package service

import (
	"battleHall/base"
	"battleHall/common"
	"battleHall/conf"
	"battleHall/dao"
	"battleHall/grpc-base/protos"
	"battleHall/model"
	"fmt"
	"google.golang.org/grpc"
	"sync"
)

var (
	lock      sync.RWMutex
	serverMap map[string]*model.ServerInfo

	playerLock sync.RWMutex
	playerMap  map[int64]string
)

func InitHallService(rpcServer *grpc.Server) {
	serverMap = make(map[string]*model.ServerInfo)
	playerMap = make(map[int64]string)

	protos.RegisterHallServiceServer(rpcServer, new(hallService))

	registService(int64(common.MAX_SERVICE_TIME))
	loadServerInfo()

	InitRpcClient()
}

func OnHallRelease() {
	for _, data := range serverMap {
		dao.ServerInfoDao.InsertOrUpdateServerInfo(data)
	}
}

func registService(ttl int64) {
	endpoint := conf.Server.EtcdServerAddr
	serviceKey := common.SERVICE_PREFIX + conf.Server.RpcServerAddr
	register := base.NewServiceRegister(endpoint, common.SERVICE_PREFIX, serviceKey, conf.Server.RpcServerAddr)
	err := register.Register(ttl)
	if err != nil {
		fmt.Printf("registService err : %+v\n", err)
	}
}

func loadServerInfo() {
	lst := dao.ServerInfoDao.LoadAll()
	for i := 0; i < len(lst); i++ {
		serverMap[lst[i].ServerAddr] = lst[i]
	}
}

func addOrUpdateServer(data *model.ServerInfo) {
	lock.Lock()
	defer lock.Unlock()
	serverMap[data.ServerAddr] = data
	dao.ServerInfoDao.InsertOrUpdateServerInfo(data)
}

func removeSceneServer(battleServerAddr string) {
	lock.Lock()
	defer lock.Unlock()
	if serv, ok := serverMap[battleServerAddr]; ok {
		for i := 0; i < len(serv.Online); i++ {
			serv.Online[i].OnlineNum = 0
		}
		serv.StartUp = false
		dao.ServerInfoDao.InsertOrUpdateServerInfo(serv)
		delete(serverMap, battleServerAddr)
	}
}

func removeGsServer(gsServer string) {
	lock.Lock()
	defer lock.Unlock()
	for _, serv := range serverMap {
		for i := 0; i < len(serv.Online); i++ {
			if serv.Online[i].ServerAddr == gsServer {
				serv.Online[i].OnlineNum = 0
				break
			}
		}
		dao.ServerInfoDao.InsertOrUpdateServerInfo(serv)
	}
}

func findServerInfo(data *model.ServerInfo) *model.ServerInfo {
	lock.RLock()
	defer lock.RUnlock()
	if serv, ok := serverMap[data.ServerAddr]; ok {
		return serv
	}
	return nil
}

// selectOneSceneClient 选择一个客户端
func selectOneSceneClient() (protos.SceneServiceClient, string) {
	lock.RLock()
	defer lock.RUnlock()
	for serverAddr, data := range serverMap {
		onlineNum := 0
		for i := 0; i < len(data.Online); i++ {
			onlineNum += data.Online[i].OnlineNum
		}
		if uint32(onlineNum) < data.OnlineLimit {
			return getSceneClient(serverAddr), serverAddr
		}
	}
	return nil, ""
}

func addScenePlayerNum(gsServer, sceneServer string) {
	lock.Lock()
	defer lock.Unlock()
	if serv, ok := serverMap[sceneServer]; ok {
		var insert = true
		for i := 0; i < len(serv.Online); i++ {
			if serv.Online[i].ServerAddr == gsServer {
				serv.Online[i].OnlineNum += 1
				insert = false
				break
			}
		}
		if insert {
			serv.Online = append(serv.Online, model.NewGsServerOnlie(gsServer))
			dao.ServerInfoDao.InsertOrUpdateServerInfo(serv)
		}
	}
}

func addScenePlayer(accountId int64, battleServer string) {
	playerLock.Lock()
	defer playerLock.Unlock()
	playerMap[accountId] = battleServer
}

func removeBattlePlayerNum(gsServer, battleServer string) {
	lock.Lock()
	defer lock.Unlock()
	if serv, ok := serverMap[battleServer]; ok {
		for i := 0; i < len(serv.Online); i++ {
			if serv.Online[i].ServerAddr == gsServer {
				serv.Online[i].OnlineNum -= 1
				if serv.Online[i].OnlineNum < 0 {
					serv.Online[i].OnlineNum = 0
				}
			}
		}
	}
}

func removePlayerInBattle(accountId int64) {
	playerLock.Lock()
	defer playerLock.Unlock()
	delete(playerMap, accountId)
}

func removeAllPlayerInScene(sceneAddr string) {
	playerLock.Lock()
	defer playerLock.Unlock()
	for key, addr := range playerMap {
		if addr == sceneAddr {
			delete(playerMap, key)
		}
	}
}

func playerInBattle(accountId int64) bool {
	playerLock.RLock()
	defer playerLock.RUnlock()

	if _, ok := playerMap[accountId]; ok {
		return true
	}
	return false
}

func getPlayerSceneServer(accountId int64) string {
	playerLock.RLock()
	defer playerLock.RUnlock()

	if serverAddr, ok := playerMap[accountId]; ok {
		return serverAddr
	}
	return ""
}

func getAllScene() []*model.ServerInfo {
	lock.RLock()
	defer lock.RUnlock()

	var ret []*model.ServerInfo
	for _, data := range serverMap {
		ret = append(ret, data)
	}
	return ret
}

func UpdateSceneLimit(sceneAddr string, onlineLimit uint32) {
	lock.Lock()
	defer lock.Unlock()

	if data, ok := serverMap[sceneAddr]; ok {
		data.OnlineLimit = onlineLimit
		dao.ServerInfoDao.InsertOrUpdateServerInfo(data)
	}
}
