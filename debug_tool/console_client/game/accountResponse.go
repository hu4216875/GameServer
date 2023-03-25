package game

import (
	"client/msg"
	"fmt"
	"github.com/golang/protobuf/proto"
	"time"
)

func registAccountResponse() {
	registResponse(msg.MsgId_ID_ResponseRegist, responseRegist)
	registResponse(msg.MsgId_ID_ResponseLogin, responseLogin)
	registResponse(msg.MsgId_ID_ResponseKickOut, responseKnickOut)
}

func responseRegist(data []byte, length int) {
	temp := &msg.ResponseRegist{}
	if err := proto.Unmarshal(data[6:length], temp); err == nil {
		fmt.Printf("recv len:%v msgId:%v ", length, responseRegist)
	}
	fmt.Printf("ResponseRegist %v", temp.Result)
}

func responseLogin(data []byte, length int) {
	temp := &msg.ResponseLogin{}
	if err := proto.Unmarshal(data[6:length], temp); err == nil {
		fmt.Printf("recv len:%v msgId:%v ", length, msg.MsgId_ID_ResponseLogin)
	}
	fmt.Printf("ResponseLogin %v\n", temp.Result)
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				sendClientHert()
			}
		}
	}()
}

func responseKnickOut(data []byte, length int) {
	temp := &msg.ResponseKickOut{}
	if err := proto.Unmarshal(data[6:length], temp); err == nil {
		fmt.Printf("recv len:%v msgId:%v ", length, msg.MsgId_ID_ResponseLogin)
	}
	fmt.Printf("ResponseKickOut %v\n", temp.Result)
}
