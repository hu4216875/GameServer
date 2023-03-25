package model

import "time"

type AccountItem struct {
	AccountId int64
	Items     []*Item
}

type Item struct {
	Id         uint32
	Num        uint32
	LimitDate  uint32
	CreateTime uint32
	UpdateTime uint32
}

func NewItem(id, num, limitDate uint32) *Item {
	curTime := uint32(time.Now().Unix())
	return &Item{
		Id:         id,
		Num:        num,
		LimitDate:  limitDate,
		CreateTime: curTime,
	}
}

func NewAccountItem(accountId int64) *AccountItem {
	ret := &AccountItem{
		AccountId: accountId,
	}
	ret.Items = make([]*Item, 0, 0)
	return ret
}
