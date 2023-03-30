package model

import (
	"server/conf"
	"time"
)

// Account 账户表
type Account struct {
	UserId     string
	Pwd        string
	AccountId  int64
	Nick       string
	ServerId   uint32
	CreateTime uint32
	UpdateTime uint32
	Forbidden  bool // 是否禁止登录
}

func NewAccount(userId string, pwd string, accountId int64) *Account {
	curTime := uint32(time.Now().Unix())
	return &Account{
		UserId:     userId,
		Pwd:        pwd,
		AccountId:  accountId,
		ServerId:   conf.Server.ServerId,
		CreateTime: curTime,
	}
}
