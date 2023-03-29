package model

type PlayerData struct {
	AccountId    int64
	GsServerAddr string
}

func NewPlayerData(accountId int64, gsServerAddr string) *PlayerData {
	return &PlayerData{
		AccountId:    accountId,
		GsServerAddr: gsServerAddr,
	}
}
