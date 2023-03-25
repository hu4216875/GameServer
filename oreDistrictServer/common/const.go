package common

import "time"

type OreErr int32

var (
	Ore_DB_Name           = "global"
	Ore_DB_Collection     = "oreDistrict"
	DB_OP_TIME_OUT        = 5 * time.Second
	Ore_DB_Log_Collection = "oreDistrictLog"
	UINT32_MAX            = 4294967295
)

const (
	OreErr_None OreErr = iota
	OreErr_NO_RESOURCE
	OreErr_PLAYER_EXIST
	OreErr_PLYAER_NOT_EXIST = 3
	OreErr_ORE_NOT_EXIST    = 4
)
