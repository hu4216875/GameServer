package template

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
	"reflect"
	"server/conf"
)

var (
	itemTemplate       ItemTemplate
	systemItemTemplate SystemItemTemplate
)

func readRf(st interface{}) *recordfile.RecordFile {
	rf, err := recordfile.New(st)
	if err != nil {
		log.Fatal("readRf err:%v", err)
	}
	fn := reflect.TypeOf(st).Name() + ".csv"
	err = rf.Read(conf.Server.GameDataPath + fn)
	if err != nil {
		log.Fatal("%v: %v", fn, err)
	}
	return rf
}

func LoadTempalte() {
	itemTemplate.load()
	systemItemTemplate.load()

	itemTemplate.init()
	systemItemTemplate.init()
}

func GetItemTemplate() *ItemTemplate {
	return &itemTemplate
}

func GetSystemItemTemplate() *SystemItemTemplate {
	return &systemItemTemplate
}
