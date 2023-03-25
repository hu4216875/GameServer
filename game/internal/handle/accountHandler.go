package handle

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/game/internal/common"
	"server/game/internal/dao"
	"server/game/internal/service"
	"server/msg"
	"server/publicconst"
	"server/util"
	"time"
)

func RpcRegist(args []interface{}) {
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

// RpcLogin 登录
func RpcLogin(args []interface{}) {
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
	var userData = common.PlayerMgr.FindPlayerData(accountId)
	if userData != nil {
		userData.UpdateTime = uint32(time.Now().Unix())
		userData.PlayerAgent = agent
	} else {
		userData = common.NewPlayerData(userId, agent)
		userData.AccountInfo = dao.AccountDao.GetAccount(accountId)
		if userData.AccountInfo == nil {
			log.Error("AccountId:% accountData is null", accountId)
			return
		}
		common.PlayerMgr.AddPlayerData(userData)
	}
	dao.AccountDao.UpdateAccountLogin(accountId)
	userData.State = publicconst.Online
	agent.SetUserData(userData)

	res := &msg.ResponseLogin{
		Result: int32(msg.ErrCode_SUCC),
	}
	agent.WriteMsg(res)
	log.Debug("rpcLogin userId:%v login succ", userId)
}

// RequestClientHeartHandle 处理客户端心跳
func RequestClientHeartHandle(args interface{}, playerData *common.PlayerData) {
	playerData.UpdateTime = util.GetCurTime()

	service.ServMgr.GetAccountService().OnHeart(playerData.AccountInfo.AccountId)
	retMsg := &msg.ResponseClientHert{Result: int32(msg.ErrCode_SUCC)}
	playerData.PlayerAgent.WriteMsg(retMsg)
}

// RequestLogoutHandle 客户端退出
func RequestLogoutHandle(args interface{}, playerData *common.PlayerData) {
	res := &msg.ResponseLogout{
		Result: int32(msg.ErrCode_SUCC),
	}
	playerData.PlayerAgent.WriteMsg(res)
	service.ServMgr.GetAccountService().OnClose(playerData)
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
