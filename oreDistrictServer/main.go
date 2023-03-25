package main

import (
	lconf "github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
	"google.golang.org/grpc"
	"net"
	"oreDistrictServer/conf"
	"oreDistrictServer/db"
	"oreDistrictServer/service"
	"oreDistrictServer/template"
)

var logger *log.Logger

func initLog() {
	// logger
	if conf.Server.LogLevel != "" {
		var err error
		logger, err = log.New(conf.Server.LogLevel, conf.Server.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
	}
}

func main() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	initLog()
	defer logger.Close()

	db.OnInit()
	template.LoadTempalte()

	rpcs := grpc.NewServer()
	service.RegistOreService(rpcs)
	service.InitOreService()

	defer db.OnDestroy()

	lis, err := net.Listen("tcp", conf.Server.RpcServerAddr)
	log.Debug("oreDistrict server start addr:%v", conf.Server.RpcServerAddr)
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)
}
