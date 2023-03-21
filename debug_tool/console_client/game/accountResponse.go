package game

import (
	"client/msg"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func registAccountResponse() {
	registResponse(msg.MsgId_ID_ResponseRegist, responseRegist)
	registResponse(msg.MsgId_ID_ResponseLogin, responseLogin)
	registResponse(msg.MsgId_ID_ResponseKickOut, responseKnickOut)
}

func responseRegist(data []byte, len int) {
	temp := &msg.ResponseRegist{}
	if err := proto.Unmarshal(data[6:len], temp); err == nil {
		fmt.Printf("recv len:%v msgId:%v ", len, responseRegist)
	}
	fmt.Printf("ResponseRegist %v", temp.Result)
}

func responseLogin(data []byte, len int) {
	temp := &msg.ResponseLogin{}
	if err := proto.Unmarshal(data[6:len], temp); err == nil {
		fmt.Printf("recv len:%v msgId:%v ", len, msg.MsgId_ID_ResponseLogin)
	}
	fmt.Printf("ResponseLogin %v\n", temp.Result)
}

func responseKnickOut(data []byte, len int) {
	temp := &msg.ResponseKickOut{}
	if err := proto.Unmarshal(data[6:len], temp); err == nil {
		fmt.Printf("recv len:%v msgId:%v ", len, msg.MsgId_ID_ResponseLogin)
	}
	fmt.Printf("ResponseKickOut %v\n", temp.Result)
}
