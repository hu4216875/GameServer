package model

import "time"

type ServerInfo struct {
	ServerId   uint32
	RegistNum  uint32
	UpdateTime uint32
	CreateTime uint32
	LimitNum   uint32
}

func NewServerInfo(serverId uint32) *ServerInfo {
	curTime := uint32(time.Now().Unix())
	return &ServerInfo{
		ServerId:   serverId,
		UpdateTime: curTime,
		CreateTime: curTime,
	}
}
