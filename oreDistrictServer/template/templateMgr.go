package template

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
	"oreDistrictServer/conf"
	"reflect"
)

var (
	oreTemplate OreDistrictTemplate
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
	oreTemplate.load()
}

func GetOreTempalte() *OreDistrictTemplate {
	return &oreTemplate
}
