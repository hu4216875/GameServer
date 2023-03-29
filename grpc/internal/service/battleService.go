package service

import (
	"context"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/grpc-base/grpc-base/protos"
	"server/msg"
	"time"
)

func RequestEnterBattle(accountId int64) []interface{} {
	req := protos.RequestEnterHall{AccountId: accountId, ServerAddr: conf.Server.RpcServer}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cli := getBattleHall()
	if cli == nil {
		log.Error("RequestEnterBattle accountId:%v hall is null", accountId)
		return nil
	}
	res, err := cli.EnterHall(ctx, &req)
	if err != nil {
		log.Error("RequestEnterBattle err:%v", err)
		return nil
	}
	var ret []interface{}

	ret = append(ret, res.Result)
	if err != nil {
		log.Error("RequestEnterBattle account:%v err:%v", accountId, err)
		return ret
	}
	if res.Result == int32(msg.ErrCode_SUCC) {
		ret = append(ret, res.SceneAddr)
	}
	return ret
}

func RequestLeaveHall(accountId int64) interface{} {
	req := protos.RequestLeaveHall{AccountId: accountId, ServerAddr: conf.Server.RpcServer}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cli := getBattleHall()
	if cli == nil {
		log.Error("RequestLeaveHall accountId:%v hall is null", accountId)
		return nil
	}
	res, err := cli.LeaveHall(ctx, &req)
	if err != nil {
		log.Error("RequestLeaveHall err:%v\n", err)
		return msg.ErrCode_SYSTEM_ERROR
	}
	return res.Result
}
