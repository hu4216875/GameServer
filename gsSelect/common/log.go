package common

import (
	lconf "github.com/name5566/leaf/conf"
	"github.com/name5566/leaf/log"
	"gsSelect/conf"
)

var logger *log.Logger

func InitLog() {
	lconf.LogLevel = conf.Server.LogLevel
	lconf.LogPath = conf.Server.LogPath
	lconf.LogFlag = conf.LogFlag
	if conf.Server.LogLevel != "" {
		var err error
		logger, err = log.New(conf.Server.LogLevel, conf.Server.LogPath, conf.LogFlag)
		if err != nil {
			panic(err)
		}
		log.Export(logger)
	}
}

func ReleaseLog() {
	if logger != nil {
		logger.Close()
	}
}
