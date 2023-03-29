package service

import (
	"context"
	"github.com/name5566/leaf/log"
	"server/conf"
	"server/grpc-base/grpc-base/protos"
	"server/msg"
	"server/publicconst"
	"server/util"
	"time"
)

var (
	total           uint32 // 总量
	endTime         uint32 // 结束时间
	updateTotalTime uint32 // 更新总量的时间
)

func SetOreInfo(oreTotal, oreEndTime uint32) {
	total = oreTotal
	endTime = oreEndTime
	updateTotalTime = util.GetCurTime()
}

func GetOreTotal(oreId uint32) uint32 {
	oreClient := GetOreRpcClient()
	if oreClient == nil {
		log.Error("GetOreTotal oreClient is nil")
		return 0
	}
	curTime := uint32(time.Now().Unix())
	if int(curTime-updateTotalTime) < publicconst.REFRESH_ORE_INTEVAL {
		return total
	}

	req := protos.RequestOreTotal{OreId: oreId}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := oreClient.GetOreTotal(ctx, &req)
	if err != nil {
		log.Error("GetOreTotal oreId:%v err:%v", oreId, err)
	}
	total = res.Total
	endTime = res.EndTime
	updateTotalTime = curTime
	return total
}

// StartOre 开始挖矿
func StartOre(accountId int64, oreId, speed uint32) msg.ErrCode {
	oreClient := GetOreRpcClient()
	if oreClient == nil {
		log.Error("StartOre oreClient is nil")
		return 0
	}

	req := protos.RequestAddOrePlayer{
		OreId:     oreId,
		AccountId: accountId,
		ServerId:  conf.Server.ServerId,
		Speed:     speed,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := oreClient.AddOrePlayer(ctx, &req)
	if err != nil {
		log.Error("StartOre oreId:%v err:%v", oreId, err)
	}
	if res.Result == 0 {
		total = res.Total
		endTime = res.EndTime
		updateTotalTime = util.GetCurTime()
		return msg.ErrCode_SUCC
	} else if res.Result == 1 {
		return msg.ErrCode_NO_ORE_RESOURCE
	} else if res.Result == 2 {
		return msg.ErrCode_HAS_START_ORE
	}
	return msg.ErrCode_SYSTEM_ERROR
}

// EndOre 结束挖矿
func EndOre(accountId int64, oreId uint32) uint32 {
	oreClient := GetOreRpcClient()
	if oreClient == nil {
		log.Error("EndOre oreClient is nil")
		return 0
	}

	req := protos.RequestSettleOrePlayer{
		OreId:     oreId,
		AccountId: accountId,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := oreClient.SettlePlayer(ctx, &req)
	if err != nil {
		log.Error("AccountId:%v EndOre oreId:%v err:%v", accountId, oreId, err)
		return 0
	}
	if res.Result == 0 {
		total = res.Total
		endTime = res.EndTime
		updateTotalTime = util.GetCurTime()
		return res.Num
	} else {
		log.Error("Account:%v EndOre:%v result:%v", accountId, res.Result)
	}
	return 0
}

// UpgradeOreSpeed 升级挖矿速度
func UpgradeOreSpeed(accountId int64, oreId, newSpeed uint32) []interface{} {
	oreClient := GetOreRpcClient()
	if oreClient == nil {
		log.Error("UpgradeOreSpeed oreClient is nil")
		return nil
	}

	var ret []interface{}
	req := protos.RequestUpdateOrePlayer{
		OreId:     oreId,
		AccountId: accountId,
		Speed:     newSpeed,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := oreClient.UpdateOrePlayer(ctx, &req)
	if err != nil {
		log.Error("AccountId:%v EndOre oreId:%v err:%v", accountId, oreId, err)
		return nil
	}

	retErr := msg.ErrCode_SUCC
	if res.Result == 0 {
		total = res.Total
		endTime = res.EndTime
		updateTotalTime = util.GetCurTime()
	} else if res.Result == 1 {
		retErr = msg.ErrCode_NO_ORE_RESOURCE
	} else if res.Result == 2 {
		retErr = msg.ErrCode_HAS_START_ORE
	} else {
		retErr = msg.ErrCode_SYSTEM_ERROR
	}
	ret = append(ret, retErr)
	ret = append(ret, res.Num)
	return ret
}

func GetEndTime() uint32 {
	return endTime
}
