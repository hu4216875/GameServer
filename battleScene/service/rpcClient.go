package service

import (
	"battleScene/base"
	"battleScene/common"
	"battleScene/conf"
	"battleScene/grpc-base/protos"
	"fmt"
	"github.com/name5566/leaf/log"
	"google.golang.org/grpc"
	"strings"
	"sync"
)

var (
	gsServerLock sync.RWMutex
	gsServerMap  map[string]protos.GsSceneServiceClient
)

func InitRpcClient() {
	gsServerMap = make(map[string]protos.GsSceneServiceClient)
	discover := base.NewServiceDiscover(conf.Server.EtcdServerAddr)
	discover.ListFunc = listRpcServer
	discover.UpdateFunc = updateRpcServer
	discover.RemoveFunc = removeRpcServer
	if err := discover.DiscoverService("/server"); err != nil {
		log.Error("InitRpcClient err:%v", err)
	}
}

func listRpcServer(key, value string) {
	if strings.Contains(key, common.GS_SERVER_PREFIX) {
		connGsRpc(value)
	}
	log.Debug("listRpcServer key:%v, value:%v ", key, value)
}

func updateRpcServer(key, value string) {
	if strings.Contains(key, common.GS_SERVER_PREFIX) {
		connGsRpc(value)
	}
	fmt.Println("update ", key, " value:", value)
}

func removeRpcServer(key string) {
	pos := strings.LastIndex(key, "/")
	var servAddr string
	if pos > 0 {
		servAddr = key[pos+1:]
		servAddr = strings.Trim(servAddr, "")
	}

	if len(servAddr) == 0 {
		return
	}
	if strings.Contains(key, common.GS_SERVER_PREFIX) {
		gsServerLock.Lock()
		defer gsServerLock.Unlock()
		RemoveGsPlayer(servAddr)
		if _, ok := gsServerMap[servAddr]; ok {
			delete(gsServerMap, servAddr)
		}
	}
	log.Error("removeRpcServer key:%v", key)
}

func connGsRpc(serverAddr string) {
	gsClient, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("connBattleRpc err:%v", err))
	}

	gsServerLock.Lock()
	defer gsServerLock.Unlock()
	gsServerMap[serverAddr] = gsClient
}

func getGsClient(gsAddr string) protos.GsSceneServiceClient {
	gsServerLock.RLock()
	defer gsServerLock.RUnlock()

	if data, ok := gsServerMap[gsAddr]; ok {
		return data
	}
	return nil
}
