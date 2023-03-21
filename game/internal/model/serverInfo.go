package model

import "time"

type ServerInfo struct {
	ServerAddr string
	RegistNum  int64
	UpdateTime uint32
	CreateTime uint32
}

func NewServerInfo(serverAddr string) *ServerInfo {
	curTime := uint32(time.Now().Unix())
	return &ServerInfo{
		ServerAddr: serverAddr,
		UpdateTime: curTime,
		CreateTime: curTime,
	}
}
