package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var LogFlag = log.LstdFlags

var Server struct {
	LogLevel       string
	LogPath        string
	RpcServerAddr  string
	GameDataPath   string
	MaxPlayerNum   int32
	EtcdServerAddr []string
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
