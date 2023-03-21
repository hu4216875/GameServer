package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

var Server struct {
	ServerId             int
	LogLevel             string
	LogPath              string
	WSAddr               string
	CertFile             string
	KeyFile              string
	TCPAddr              string
	MaxConnNum           int
	ConsolePort          int
	ProfilePath          string
	MaxLocalMongoDBConn  int
	LocalMongoDBUrl      string
	MaxGlobalMongoDBConn int
	GlobalMongoDBUrl     string
	GameDataPath         string
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
