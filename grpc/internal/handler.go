package internal

import (
	"reflect"
	"server/grpc/internal/service"
)

func init() {
	skeleton.RegisterChanRPC("GetOreTotal", getOreTotal)
	skeleton.RegisterChanRPC("StartOre", startOre)
	skeleton.RegisterChanRPC("EndOre", endOre)
	skeleton.RegisterChanRPC("UpgradeOreSpeed", upgradeOreSpeed)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

// getOreTotal 登录
func getOreTotal(args []interface{}) interface{} {
	oreId := args[0].(uint32)
	return service.GetOreTotal(oreId)
}

// startOre 开始挖矿
func startOre(args []interface{}) interface{} {
	oreId := args[0].(uint32)
	accountId := args[1].(int64)
	speed := args[2].(uint32)
	return service.StartOre(accountId, oreId, speed)
}

// EndOre 结算挖矿
func endOre(args []interface{}) interface{} {
	oreId := args[0].(uint32)
	accountId := args[1].(int64)
	return service.EndOre(accountId, oreId)
}

func upgradeOreSpeed(args []interface{}) []interface{} {
	accountId := args[0].(int64)
	oreId := args[1].(uint32)
	newSpeed := args[2].(uint32)
	return service.UpgradeOreSpeed(accountId, oreId, newSpeed)
}
