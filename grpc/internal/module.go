package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"google.golang.org/grpc"
	"net"
	"server/base"
	"server/conf"
	"server/grpc/internal/service"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer

	oreServerClient *grpc.ClientConn
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	service.InitRpcClient()

	skeleton.Go(func() {
		m.startRpcServer()
	}, func() {

	})
}

func (m *Module) OnDestroy() {
	if oreServerClient != nil {
		oreServerClient.Close()
	}
}

func (m *Module) startRpcServer() {
	rpcs := grpc.NewServer()
	service.RegistRpcServer(rpcs)
	lis, err := net.Listen("tcp", conf.Server.RpcServer)
	log.Debug("rpc server start addr %v", conf.Server.RpcServer)
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)
}

func (m *Module) SetSceneCloseFunc(sceneClose func(sceneAddr string)) {
	service.SetSceneCloseFunc(sceneClose)
}
