package model

import "time"

type Account struct {
	UserId     string
	AccountId  int64
	Nick       string
	LoginTime  uint32
	LogoutTime uint32
	CreateTime uint32
	UpdateTime uint32
	OreInfo    *OreInfo
}

// OreInfo 矿洞信息
type OreInfo struct {
	OreId     uint32
	StartTime uint32
	Speed     uint32
}

func NewAccount(userId string, accountId int64) *Account {
	curTime := uint32(time.Now().Unix())
	return &Account{
		UserId:     userId,
		AccountId:  accountId,
		LoginTime:  curTime,
		CreateTime: curTime,
		OreInfo:    &OreInfo{},
	}
}
