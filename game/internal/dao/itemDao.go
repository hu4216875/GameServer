package dao

import (
	"server/game/internal/model"
	"server/msg"
)

var (
	ItemDao = &itemDao{}
)

type itemDao struct {
}

func (a *itemDao) LoadItem(userId string, accountId int64) (msg.ErrCode, []*model.Item) {
	return msg.ErrCode_SUCC, nil
}

func (a *itemDao) AddItem(userId string, accountId int64) msg.ErrCode {
	return msg.ErrCode_SUCC
}
