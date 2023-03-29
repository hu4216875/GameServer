package service

import (
	"battleScene/conf"
	"battleScene/grpc-base/protos"
	"battleScene/model"
	"context"
)

type sceneRpcService struct {
	protos.UnimplementedSceneServiceServer
}

func (s *sceneRpcService) EnterScene(ctx context.Context, req *protos.RequestEnterScene) (*protos.ResponseEnterScene, error) {
	playerData := model.NewPlayerData(req.AccountId, req.GsServerAddr)
	AddPlayerData(playerData)
	enter(playerData)
	return &protos.ResponseEnterScene{Result: int32(protos.Grpc_ErrCode_SUCC)}, nil
}
func (s *sceneRpcService) LeaveScene(ctx context.Context, req *protos.RequestLeaveScene) (*protos.ResponseLeaveScene, error) {
	leave(req.AccountId)
	RemovePlayerData(req.AccountId)
	return &protos.ResponseLeaveScene{}, nil
}

func (s *sceneRpcService) GetSceneInfo(ctx context.Context, req *protos.RequestServerInfo) (*protos.ResponseServerInfo, error) {
	return &protos.ResponseServerInfo{PlayerMaxLimit: conf.Server.MaxPlayerNum}, nil
}
