package internal

import (
	"github.com/name5566/leaf/module"
	"log"
	"server/base"
	"server/conf"
	"server/util"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer

	snowWorker *util.SnowWorker
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton

	var err error
	if snowWorker, err = util.NewSnowWorker(int64(conf.Server.ServerId)); err != nil {
		log.Fatalf("game module init err:%v", err)
	}
}

func (m *Module) OnDestroy() {

}
