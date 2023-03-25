package game

import (
	"client/msg"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

type CommandItem struct {
	ItemId uint32 `json:"itemId"`
	Num    uint32 `json:"num"`
}

func SendGmAddItem() {
	var items []*CommandItem
	items = append(items, &CommandItem{
		ItemId: 1,
		Num:    1000,
	})
	items = append(items, &CommandItem{
		ItemId: 2,
		Num:    1,
	})
	items = append(items, &CommandItem{
		ItemId: 3,
		Num:    3,
	})

	content, err := json.Marshal(items)
	if err != nil {
		fmt.Printf("SendGmAddItem err:%v", err)
	}

	msgTemp := &msg.RequestGMCommand{}
	msgTemp.CommandId = int32(msg.GMCommand_AddItem)
	msgTemp.Content = string(content)
	// 进行编码
	data, err := proto.Marshal(msgTemp)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestGMCommand), data)
}

func SendGmReload() {
	var files []string
	files = append(files, "system")
	content, err := json.Marshal(files)
	if err != nil {
		fmt.Printf("SendGmReload err:%v", err)
	}

	msgTemp := &msg.RequestGMCommand{}
	msgTemp.CommandId = int32(msg.GMCommand_ReloadConfig)
	msgTemp.Content = string(content)
	// 进行编码
	data, err := proto.Marshal(msgTemp)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestGMCommand), data)
}
