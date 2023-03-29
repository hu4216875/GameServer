package service

import (
	"battleHall/grpc-base/protos"
	"context"
	"github.com/name5566/leaf/log"
	"time"
)

type hallService struct {
	protos.UnimplementedHallServiceServer
}

func (h *hallService) EnterHall(ctx context.Context, req *protos.RequestEnterHall) (*protos.ResponseEnterHall, error) {
	res := &protos.ResponseEnterHall{Result: int32(protos.Grpc_ErrCode_SUCC)}
	if playerInBattle(req.AccountId) {
		res.Result = int32(protos.Grpc_ErrCode_PALYER_IN_BATTLE)
		return res, nil
	}

	sceneClient, sceneServerAddr := selectOneSceneClient()
	if sceneClient == nil {
		res.Result = int32(protos.Grpc_ErrCode_SCENE_SERVER_FULL)
		return res, nil
	}

	sceneReq := protos.RequestEnterScene{AccountId: req.AccountId, GsServerAddr: req.ServerAddr}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sceneRes, err := sceneClient.EnterScene(ctx, &sceneReq)
	if err != nil {
		log.Error("EnterBattle AccountId:%v err:%v", req.AccountId, err)
		res.Result = int32(protos.Grpc_ErrCode_SCENE_NOT_EXIST)
		return res, nil
	}
	if sceneRes.Result == int32(protos.Grpc_ErrCode_SUCC) {
		res.SceneAddr = sceneServerAddr
		addScenePlayerNum(req.ServerAddr, sceneServerAddr)
		addScenePlayer(req.AccountId, sceneServerAddr)
	}
	return res, nil
}

func (h *hallService) LeaveHall(ctx context.Context, req *protos.RequestLeaveHall) (*protos.ResponseLeaveHall, error) {
	sceneAddr := getPlayerSceneServer(req.AccountId)
	sceneServer := getSceneClient(sceneAddr)
	if sceneServer == nil {
		log.Error("LeaveBattle accountId:%v sceneAddr:%v not exist", req.AccountId, sceneAddr)
		return &protos.ResponseLeaveHall{Result: int32(protos.Grpc_ErrCode_SCENE_NOT_EXIST)}, nil
	}
	removeBattlePlayerNum(req.ServerAddr, sceneAddr)
	removePlayerInBattle(req.AccountId)

	sceneReq := protos.RequestLeaveScene{AccountId: req.AccountId}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	sceneServer.LeaveScene(ctx, &sceneReq)

	return &protos.ResponseLeaveHall{Result: int32(protos.Grpc_ErrCode_SUCC)}, nil
}
