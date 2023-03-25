package game

import (
	"client/msg"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func registCmCommandResponse() {
	registResponse(msg.MsgId_ID_ResponseGMCommand, responseGmCommand)
}

func responseGmCommand(data []byte, length int) {
	res := &msg.ResponseGMCommand{}
	if err := proto.Unmarshal(data[6:length], res); err != nil {
		fmt.Printf("recv len:%v msgId:%v:err:%v\n", length, res, err)
	}

	fmt.Printf("responseGmCommand result:%v\n", res.Result)

}
