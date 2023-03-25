package main

import (
	"github.com/name5566/leaf"
	lconf "github.com/name5566/leaf/conf"
	"server/conf"
	"server/db"
	"server/game"
	"server/gate"
	"server/grpc"
	"server/login"
	"server/template"
)

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	lconf.ConsolePort = conf.Server.ConsolePort
	lconf.ProfilePath = conf.Server.ProfilePath

	// 初始化陪标数据
	template.LoadTempalte()

	leaf.Run(
		db.Module,
		grpc.Module,
		game.Module,
		gate.Module,
		login.Module,
	)
}
