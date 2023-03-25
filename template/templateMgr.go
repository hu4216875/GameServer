package template

import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/recordfile"
	"reflect"
	"server/conf"
)

type IDataTempalte interface {
	OnInit()
	OnLoad()
	OnReload()
	OnDestory()
}

var (
	templateMap map[string]IDataTempalte
)

func init() {
	templateMap = make(map[string]IDataTempalte)
	templateMap[ITEM_FILE] = NewItemTemplate()
	templateMap[SYSTEM_FILE] = NewSystemItemTemplate()
}

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
	// 加载所有配置文件
	for _, data := range templateMap {
		data.OnLoad()
	}

	// 所有配置文件加载之后初始化
	for _, data := range templateMap {
		data.OnInit()
	}
}

func ReloadTemplate(files []string) {
	if len(files) == 0 {
		for _, data := range templateMap {
			data.OnReload()
		}
	} else {
		for i := 0; i < len(files); i++ {
			if data, ok := templateMap[files[i]]; ok {
				data.OnReload()
			}
		}
	}
}

func GetItemTemplate() *ItemTemplate {
	if data, ok := templateMap[ITEM_FILE]; ok {
		return data.(*ItemTemplate)
	}
	return nil
}

func GetSystemItemTemplate() *SystemItemTemplate {
	if data, ok := templateMap[SYSTEM_FILE]; ok {
		return data.(*SystemItemTemplate)
	}
	return nil
}
