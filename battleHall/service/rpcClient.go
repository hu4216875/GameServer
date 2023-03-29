package service

import (
	"battleHall/base"
	"battleHall/common"
	"battleHall/conf"
	"battleHall/grpc-base/protos"
	"battleHall/model"
	"context"
	"fmt"
	"github.com/name5566/leaf/log"
	"google.golang.org/grpc"
	"strings"
	"sync"
	"time"
)

var (
	clientLock  sync.RWMutex
	sceneClient = make(map[string]protos.SceneServiceClient)
)

func InitRpcClient() {
	discover := base.NewServiceDiscover(conf.Server.EtcdServerAddr)
	discover.ListFunc = listRpcServer
	discover.UpdateFunc = updateRpcServer
	discover.RemoveFunc = removeRpcServer
	if err := discover.DiscoverService("/server"); err != nil {
		log.Error("InitRpcClient err:%v", err)
	}
}

func listRpcServer(key, value string) {
	if strings.Contains(key, common.RPC_SCENE_SERVER_PREFIX) {
		connSceneRpc(value)
	}
	log.Debug("listRpcServer key:%v, value:%v ", key, value)
}

func updateRpcServer(key, value string) {
	if strings.Contains(key, common.RPC_SCENE_SERVER_PREFIX) {
		connSceneRpc(value)
	}
	fmt.Println("update ", key, " value:", value)
}

func removeRpcServer(key string) {
	clientLock.Lock()
	defer clientLock.Unlock()
	pos := strings.LastIndex(key, "/")
	var servAddr string
	if pos > 0 {
		servAddr = key[pos+1:]
		servAddr = strings.Trim(servAddr, "")
	}

	if len(servAddr) == 0 {
		return
	}
	if strings.Contains(key, common.RPC_SCENE_SERVER_PREFIX) {
		removeSceneServer(servAddr)
		removeAllPlayerInScene(servAddr)
		if _, ok := sceneClient[servAddr]; ok {
			delete(sceneClient, servAddr)
		}
	} else if strings.Contains(key, common.RPC_GS_SERVER_PREFIX) {
		removeGsServer(servAddr)
	}
	log.Error("removeRpcServer key:%v", key)
}

func connSceneRpc(sceneAddr string) {
	sceneServerClient, err := grpc.Dial(sceneAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("connSceneRpc err:%v", err))
	}
	cli := protos.NewSceneServiceClient(sceneServerClient)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err2 := cli.GetSceneInfo(ctx, &protos.RequestServerInfo{})
	if err2 != nil {
		log.Error("connSceneRpc err2:%v", err2)
		return
	}

	sceneInfo := model.NewServerInfo(sceneAddr, uint32(res.PlayerMaxLimit))
	sceneInfo.StartUp = true

	clientLock.Lock()
	defer clientLock.Unlock()
	sceneClient[sceneAddr] = cli
	addOrUpdateServer(sceneInfo)
}

func getSceneClient(serverAddr string) protos.SceneServiceClient {
	if data, ok := sceneClient[serverAddr]; ok {
		return data
	}
	return nil
}
