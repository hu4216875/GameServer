package service

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"server/game/internal/common"
	"server/msg"
	"server/publicconst"
	"server/template"
)

type GmService struct {
	IService
}

type CommandItem struct {
	ItemId uint32 `json:"itemId"`
	Num    uint32 `json:"num"`
}

func NewGmService() *GmService {
	return &GmService{}
}

func (g *GmService) OnInit() {
}

func (g *GmService) OnDestory() {

}

func (g *GmService) ProcessCommand(command int32, content string, playerData *common.PlayerData) msg.ErrCode {
	switch msg.GMCommand(command) {
	case msg.GMCommand_AddItem:
		return g.addItem(content, playerData)
	case msg.GMCommand_ReloadConfig:
		return g.reload(content)
	}
	return msg.ErrCode_SYSTEM_ERROR
}

func (g *GmService) addItem(content string, playerData *common.PlayerData) msg.ErrCode {
	var items []*CommandItem
	json.Unmarshal([]byte(content), &items)

	for i := 0; i < len(items); i++ {
		ServMgr.GetItemService().AddItem(playerData.AccountInfo.AccountId, items[i].ItemId, items[i].Num, publicconst.GMAddItem)
	}
	return msg.ErrCode_SUCC
}

func (g *GmService) reload(content string) msg.ErrCode {
	var files []string
	if err := json.Unmarshal([]byte(content), &files); err != nil {
		log.Error("GmService reload err:%v", err)
		return msg.ErrCode_SYSTEM_ERROR
	}
	template.ReloadTemplate(files)
	return msg.ErrCode_SUCC
}
