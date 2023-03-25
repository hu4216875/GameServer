package model

import (
	"oreDistrictServer/util"
	"time"
)

// OreDistrictPlayer 矿区玩家
type OreDistrictPlayer struct {
	AccountId int64
	ServerId  uint32
	Speed     uint32
	StartTime uint32
}

type OreDistrict struct {
	OreDistId  uint32 // 矿区Id
	EndTime    uint32 // 矿洞资源挖完的时间
	Total      uint32 // 当前总量
	Players    []*OreDistrictPlayer
	CreateTime uint32
	UpdateTime uint32
}

type OreDistrictLog struct {
	OreId   uint32
	Records []*OreDistrictRecord
}

type OreDistrictRecord struct {
	AccountId  int64
	Num        uint32
	CreateTime string
}

func NewOreDistirct(id, total uint32) *OreDistrict {
	ret := &OreDistrict{
		OreDistId: id,
		Total:     total,
	}
	ret.CreateTime = uint32(time.Now().Unix())
	ret.Players = make([]*OreDistrictPlayer, 0, 0)
	return ret
}

func NewOreDistirctPlayer(accountId int64, serverId, speed, startTime uint32) *OreDistrictPlayer {
	return &OreDistrictPlayer{
		AccountId: accountId,
		ServerId:  serverId,
		Speed:     speed,
		StartTime: startTime,
	}
}

func NewOreDistrictLog(oreId uint32) *OreDistrictLog {
	ret := &OreDistrictLog{
		OreId: oreId,
	}
	ret.Records = make([]*OreDistrictRecord, 0, 0)
	return ret
}

func NewOreDistrictRecord(accountId int64, num uint32) *OreDistrictRecord {
	return &OreDistrictRecord{
		AccountId:  accountId,
		Num:        num,
		CreateTime: util.GetLocalTimeStr(),
	}
}
