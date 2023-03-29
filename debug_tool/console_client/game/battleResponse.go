package game

import (
	"client/msg"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func registBattleResponse() {
	registResponse(msg.MsgId_ID_ResponseEnterBattle, responseEnter)
	registResponse(msg.MsgId_ID_ResponseLeaveBattle, responsLeave)
}

func responseEnter(data []byte, length int) {
	temp := &msg.ResponseEnterBattle{}
	if err := proto.Unmarshal(data[6:length], temp); err == nil {
		fmt.Printf("recv len:%v msgId:%v ", length, msg.MsgId_ID_ResponseLeaveBattle)
	}
	fmt.Printf("responseBattleEnter result:%v", temp.Result)
}

func responsLeave(data []byte, length int) {
	temp := &msg.ResponseLeaveBattle{}
	if err := proto.Unmarshal(data[6:length], temp); err == nil {
		fmt.Printf("recv len:%v msgId:%v ", length, msg.MsgId_ID_ResponseLeaveBattle)
	}
	fmt.Printf("responsBattleLeave %v\n", temp.Result)
}
