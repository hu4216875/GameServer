package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
	"server/base"
	"server/conf"
	"server/db"
	"server/game"
	"server/gate"
	"server/grpc"
	"server/login"
	"server/publicconst"
	"server/template"
)

func registServer() {
	serviceKey := publicconst.SERVER_PREFIX + conf.Server.TCPAddr
	register := base.NewServiceRegister(conf.Server.EtcdServer, publicconst.SERVER_PREFIX, serviceKey, conf.Server.TCPAddr)
	if err := register.Register(int64(publicconst.MAX_SERVER_TTL)); err != nil {
		log.Error("err:%v", err)
	}
}

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	// 初始化配表数据
	template.LoadTempalte()

	// 注册服务
	registServer()

	leaf.Run(
		db.Module,
		grpc.Module,
		game.Module,
		gate.Module,
		login.Module,
	)
}
