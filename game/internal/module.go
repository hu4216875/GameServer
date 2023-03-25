package internal

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"server/base"
	"server/conf"
	"server/game/internal/common"
	"server/game/internal/dao"
	"server/game/internal/service"
	"time"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	UpdateServerInfo()
	m.initHeartTicker()
	service.ServMgr.InitService()
}

func (m *Module) OnDestroy() {
	service.ServMgr.Destory()
}

func (m *Module) initHeartTicker() {
	ticker := time.NewTicker(10 * time.Second)
	go clientHeartCheck(ticker)
}

func clientHeartCheck(ticker *time.Ticker) {
	defer func() {
		if err := recover(); err != nil {
			log.Error("clientHeartCheck err:%v", err)
		}
		go clientHeartCheck(ticker)
	}()

	select {
	case <-ticker.C:
		checkClientLive()
		common.PlayerMgr.RecyclePlayerData()
	}

}

func UpdateServerInfo() {
	if dao.ServerInfoDao.ExistServerInfo(conf.Server.TCPAddr) {
		dao.ServerInfoDao.UpdateServerTime(conf.Server.TCPAddr)
	} else {
		dao.ServerInfoDao.AddServerInfo(conf.Server.TCPAddr)
	}
}

func checkClientLive() {
	offlinePlayers := common.PlayerMgr.GetOfflinePlayer()
	for i := 0; i < len(offlinePlayers); i++ {
		service.ServMgr.GetAccountService().OnClose(offlinePlayers[i])
		log.Debug("##### userId:%v offline", offlinePlayers[i].UserId)
	}
}
