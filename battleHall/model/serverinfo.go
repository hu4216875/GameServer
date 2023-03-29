package model

type ServerInfo struct {
	ServerAddr  string            `json:"serverAddr"`  //
	Online      []*GsServerOnline `json:"online"`      // 在线人数
	OnlineLimit uint32            `json:"onlineLimit"` // 在线人数上限
	StartUp     bool              `json:"StartUp"`     // 是否启动
}

type GsServerOnline struct {
	ServerAddr string `json:"serverAddr"`
	OnlineNum  int    `json:"onlineNum"`
}

func NewGsServerOnlie(gsServerAddr string) *GsServerOnline {
	return &GsServerOnline{ServerAddr: gsServerAddr, OnlineNum: 1}
}

func NewServerInfo(serverAddr string, limit uint32) *ServerInfo {
	ret := &ServerInfo{
		ServerAddr:  serverAddr,
		OnlineLimit: uint32(limit),
	}

	ret.Online = make([]*GsServerOnline, 0, 0)
	return ret
}
