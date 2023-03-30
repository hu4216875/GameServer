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
	MaxMongoDbConn uint32
	MongoDbUrl     string
	EtcdServerAddr []string
	HttpServer     string
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
