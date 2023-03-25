package game

import (
	"client/msg"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func registOreResponse() {
	registResponse(msg.MsgId_ID_ResponseStartOre, responseStartOre)
}

func responseStartOre(data []byte, length int) {
	res := &msg.ResponseStartOre{}
	if err := proto.Unmarshal(data[6:length], res); err != nil {
		fmt.Printf("recv len:%v msgId:%v err:%v \n", length, responseRegist, err)
	}
	fmt.Printf("responseStartOre result: %v\n", res.Result)
}
