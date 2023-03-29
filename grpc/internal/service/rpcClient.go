package service

import (
	"context"
	"fmt"
	"github.com/name5566/leaf/log"
	"google.golang.org/grpc"
	"server/base"
	"server/conf"
	"server/grpc-base/grpc-base/protos"
	"server/publicconst"
	"server/template"
	"strings"
	"sync"
	"time"
)

var (
	oreClient        protos.OreDistrictServiceClient
	battleHallClient protos.HallServiceClient

	sceneMapLock   sync.RWMutex
	sceneClientMap = make(map[string]protos.SceneServiceClient)

	sceneCloseFunc func(sceneAddr string)
)

func InitRpcClient() {
	discover := base.NewServiceDiscover(conf.Server.EtcdServer)
	discover.ListFunc = listRpcServer
	discover.UpdateFunc = updateRpcServer
	discover.RemoveFunc = removeRpcServer
	if err := discover.DiscoverService(publicconst.RPC_SERVER_PREFIX); err != nil {
		log.Error("initRpcClient err:%v", err)
	}
}

func GetOreRpcClient() protos.OreDistrictServiceClient {
	return oreClient
}

func listRpcServer(key, value string) {
	if strings.Contains(key, publicconst.RPC_SERVER_ORE_PREFIX) {
		connOreRpc(value)
	} else if strings.Contains(key, publicconst.RPC_SERVER_BATTLE_HALL_PREFIX) {
		connBattleHall(value)
	} else if strings.Contains(key, publicconst.RPC_SERVER_BATTLE_SCENE_PREFIX) {
		connSceneRpc(value)
	}
	fmt.Println("list ", key, " value:", value)
}

func updateRpcServer(key, value string) {
	if strings.Contains(key, publicconst.RPC_SERVER_ORE_PREFIX) {
		connOreRpc(value)
	} else if strings.Contains(key, publicconst.RPC_SERVER_BATTLE_HALL_PREFIX) {
		connBattleHall(value)
	} else if strings.Contains(key, publicconst.RPC_SERVER_BATTLE_SCENE_PREFIX) {
		connSceneRpc(value)
	}
	fmt.Println("update ", key, " value:", value)
}

func removeRpcServer(key string) {
	if strings.Contains(key, publicconst.RPC_SERVER_ORE_PREFIX) {
		oreClient = nil
	} else if strings.Contains(key, publicconst.RPC_SERVER_BATTLE_HALL_PREFIX) {
		removeBattlHall()
	} else if strings.Contains(key, publicconst.RPC_SERVER_BATTLE_SCENE_PREFIX) {
		pos := strings.LastIndex(key, "/")
		if pos >= 0 {
			strAddr := key[pos+1:]
			strAddr = strings.Trim(strAddr, " ")
			removeSceneRpc(strAddr)
			if sceneCloseFunc != nil {
				sceneCloseFunc(strAddr)
			}
		}
	}
	fmt.Println("remove ", key)
}

func connOreRpc(serverAddr string) {
	oreServerClient, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("ConnOreServer:%v", err))
	}
	oreClient = protos.NewOreDistrictServiceClient(oreServerClient)

	req := protos.RequestOreInfo{ServerId: conf.Server.ServerId, OreId: template.GetSystemItemTemplate().GetOreId()}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := oreClient.GetOreInfo(ctx, &req)
	if err != nil {
		log.Error("ConnOreServer %v", err)
	} else {
		SetOreInfo(res.Total, res.EndTime)
	}
}

func removeSceneRpc(sceneAddr string) {
	sceneMapLock.Lock()
	defer sceneMapLock.Unlock()
	delete(sceneClientMap, sceneAddr)
}

func connSceneRpc(sceneAddr string) {
	sceneServerClient, err := grpc.Dial(sceneAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("ConnOreServer:%v", err))
	}
	sceneClient := protos.NewSceneServiceClient(sceneServerClient)

	sceneMapLock.Lock()
	defer sceneMapLock.Unlock()
	sceneClientMap[sceneAddr] = sceneClient
}

func getSceneClient(sceneAddr string) protos.SceneServiceClient {
	sceneMapLock.RLock()
	sceneMapLock.RUnlock()
	if data, ok := sceneClientMap[sceneAddr]; ok {
		return data
	}
	return nil
}

func connBattleHall(hallAddr string) {
	hallServerClient, err := grpc.Dial(hallAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(fmt.Sprintf("ConnOreServer:%v", err))
	}
	battleHallClient = protos.NewHallServiceClient(hallServerClient)
}

func removeBattlHall() {
	battleHallClient = nil
}

func getBattleHall() protos.HallServiceClient {
	return battleHallClient
}

func SetSceneCloseFunc(f func(sceneAddr string)) {
	sceneCloseFunc = f
}
