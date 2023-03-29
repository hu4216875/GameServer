package main

import (
	"battleHall/common"
	"battleHall/conf"
	"battleHall/db"
	"battleHall/service"
	"github.com/name5566/leaf/log"
	"google.golang.org/grpc"
	"net"
)

func main() {
	common.InitLog()
	defer common.ReleaseLog()
	db.OnInit()
	defer db.OnRelease()

	rpcs := grpc.NewServer()
	service.InitHallService(rpcs)
	defer service.OnHallRelease()

	// 开启http服务
	service.InitHttpService()

	lis, err := net.Listen("tcp", conf.Server.RpcServerAddr)
	log.Debug("battleHall server start addr:%v", conf.Server.RpcServerAddr)
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)
}
