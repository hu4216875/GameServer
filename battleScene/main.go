package main

import (
	"battleScene/common"
	"battleScene/conf"
	"battleScene/service"
	"github.com/name5566/leaf/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	common.InitLog()
	defer common.ReleaseLog()

	rpcs := grpc.NewServer()
	service.InitSceneService(rpcs)
	service.InitRpcClient()
	defer service.OnSceneRelease()

	lis, err := net.Listen("tcp", conf.Server.RpcServerAddr)
	log.Debug("battleScene server start addr:%v", conf.Server.RpcServerAddr)
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)
}
