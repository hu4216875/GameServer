package service

import (
	"battleScene/base"
	"battleScene/common"
	"battleScene/conf"
	"battleScene/grpc-base/protos"
	"battleScene/model"
	"fmt"
	"github.com/name5566/leaf/log"
	"google.golang.org/grpc"
)

func InitSceneService(grpc *grpc.Server) {
	registService(int64(common.MAX_SERVICE_TIME))
	protos.RegisterSceneServiceServer(grpc, new(sceneRpcService))
}

func OnSceneRelease() {

}

func registService(ttl int64) {
	endpoint := conf.Server.EtcdServerAddr
	serviceKey := common.SCENE_SERVICE_PREFIX + conf.Server.RpcServerAddr
	register := base.NewServiceRegister(endpoint, common.SCENE_SERVICE_PREFIX, serviceKey, conf.Server.RpcServerAddr)
	err := register.Register(ttl)
	if err != nil {
		fmt.Printf("registService err : %+v\n", err)
	}
}

func enter(player *model.PlayerData) {
	log.Debug("enter accountId:%v\n", player.AccountId)
}

func leave(accountId int64) {
	log.Debug("leave accountId:%v\n", accountId)
}
