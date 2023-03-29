package common

import "time"

type OreErr int32

var (
	ORE_DB_NAME           = "global"
	Ore_DB_COLLECTION     = "oreDistrict"
	DB_OP_TIME_OUT        = 5 * time.Second
	Ore_DB_LOG_COLLECTION = "oreDistrictLog"
	UINT32_MAX            = 4294967295
	MAX_SERVICE_TIME      = 10 // 续约时间
	SERVICE_PREFIX        = "/server/rpcServer/oreServer/"
)

const (
	OreErr_None OreErr = iota
	OreErr_NO_RESOURCE
	OreErr_PLAYER_EXIST
	OreErr_PLYAER_NOT_EXIST = 3
	OreErr_ORE_NOT_EXIST    = 4
)
