package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/game"
	"server/login/internal/dao"
	"server/login/internal/model"
	"server/msg"
	"server/publicconst"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handler(&msg.RequestRegist{}, handleRegist)
	handler(&msg.RequestLogin{}, handleLoginReq)
}

type HandleFunc func(args []interface{})

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

// handleRegist 处理注册
func handleRegist(args []interface{}) {
	m := args[0].(*msg.RequestRegist)
	agent := args[1].(gate.Agent)
	ret, account := doRegist(m)
	if ret != msg.ErrCode_SUCC {
		res := &msg.ResponseRegist{
			Result: int32(ret),
		}
		agent.WriteMsg(res)
	} else {
		// game rpc regist
		skeleton.Go(func() {
			if err := game.ChanRPC.Call0("Regist", account.UserId, account.AccountId, agent); err != nil {
				log.Error("handleRegist err:%v", err)
			}
		}, func() {
		})
	}
}

func doRegist(req *msg.RequestRegist) (msg.ErrCode, *model.Account) {
	if len(req.UserId) == 0 {
		return msg.ErrCode_USERID_EMPTY, nil
	}

	if len(req.UserId) > publicconst.MAX_USERID_LEN {
		return msg.ErrCode_USERID_OVER_MAX_LEN, nil
	}

	return dao.AccountDao.Regist(req.UserId, req.Passwd, snowWorker.NextId())
}

// handleLoginReq 处理登录
func handleLoginReq(args []interface{}) {
	m := args[0].(*msg.RequestLogin)
	agent := args[1].(gate.Agent)
	log.Debug("handleLoginReq userid:%v start login", m.UserId)
	res := &msg.ResponseLogin{}

	playerData := game.GetPlayerData(m.UserId)
	if playerData != nil && playerData.PlayerAgent == agent {
		if playerData.State == publicconst.Logining {
			res.Result = int32(msg.ErrCode_ISLOGINING)
			agent.WriteMsg(res)
			return
		}
	}

	// 开启携程第三方验证
	skeleton.Go(func() {
		// 账户不存在
		account := dao.AccountDao.FindAccount(m.UserId)
		var result = int32(msg.ErrCode_SUCC)
		if account == nil {
			res.Result = int32(msg.ErrCode_USER_ID_EXIST)
		} else if account.Forbidden {
			res.Result = int32(msg.ErrCode_FORBIDDEN_USER)
		}

		// 通过了校验 如果其他地方有登录则踢下线
		if playerData != nil && playerData.State != publicconst.Offline {
			kickMsg := &msg.ResponseKickOut{Result: int32(msg.ErrCode_OTHER_LOGIN)}
			playerData.PlayerAgent.WriteMsg(kickMsg)
			playerData.PlayerAgent.Close()
		}

		if result != int32(msg.ErrCode_SUCC) {
			agent.WriteMsg(res)
			return
		}
		game.ChanRPC.Call0("Login", m.UserId, account.AccountId, agent)
	}, func() {
		//	log.Debug("handleLoginReq user:%v succ", m.Username)
	})
}
