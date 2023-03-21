package model

import "time"

type Item struct {
	Id         uint32
	Num        uint32
	LimitDate  uint32
	CreateTime uint32
	UpdateTime uint32
}

func NewItem(id, num uint32) *Item {
	curTime := uint32(time.Now().Unix())
	return &Item{
		Id:         id,
		Num:        num,
		CreateTime: curTime,
	}
}
