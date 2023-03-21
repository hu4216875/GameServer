package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/game/internal/common"
	"server/game/internal/dao"
	"server/msg"
	"server/publicconst"
	"time"
)

func rpcRegist(args []interface{}) {
	userId := args[0].(string)
	accountId := args[1].(int64)
	agent := args[2].(gate.Agent)

	err := dao.AccountDao.AddAccount(userId, accountId)
	res := &msg.ResponseRegist{
		Result: int32(err),
	}

	// 注册成功的处理
	if err == msg.ErrCode_SUCC {
		// 更新注册数
		dao.ServerInfoDao.UdpateRegistNum(conf.Server.TCPAddr)
		log.Debug("userId:%v, accountId:%v rpcRegist succ", userId, accountId)
	}
	agent.WriteMsg(res)
}

// rpcLogin 登录
func rpcLogin(args []interface{}) {
	userId := args[0].(string)
	accountId := args[1].(int64)
	agent := args[2].(gate.Agent)

	// 登录中
	if checkLogining(agent) {
		res := &msg.ResponseLogin{
			Result: int32(msg.ErrCode_ISLOGINING),
		}
		agent.WriteMsg(res)
		return
	}

	// 设置userdata
	var userData = common.PlayerMgr.FindPlayerData(userId)
	if userData != nil {
		userData.UpdateTime = uint32(time.Now().Unix())
		userData.PlayerAgent = agent
	} else {
		userData = common.NewPlayerData(userId, accountId, agent)
		common.PlayerMgr.AddPlayerData(userData)
	}
	userData.State = publicconst.Online
	agent.SetUserData(userData)

	res := &msg.ResponseLogin{
		Result: int32(msg.ErrCode_SUCC),
	}
	agent.WriteMsg(res)
	log.Debug("rpcLogin userId:%v login succ", userId)
}

// rpcLogout 退出
func rpcLogout(args []interface{}) {
	username := args[0].(string)
	log.Debug("rpcLogout username:%v", username)
}

func checkLogining(agent gate.Agent) bool {
	userData := agent.UserData()
	if userData != nil {
		if playerData := userData.(*common.PlayerData); playerData != nil {
			if playerData.State == publicconst.Logining {
				return true
			}
		}
	}
	return false
}
