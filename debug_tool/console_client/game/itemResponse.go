package game

import (
	"client/msg"
	"fmt"
	"github.com/golang/protobuf/proto"
)

func registItemResponse() {
	registResponse(msg.MsgId_ID_ResponseLoadItem, responseLoadItem)
	registResponse(msg.MsgId_ID_NotifyUpdateItem, responseUpdateItem)
}

func responseLoadItem(data []byte, length int) {
	res := &msg.ResponseLoadItem{}
	if err := proto.Unmarshal(data[6:length], res); err != nil {
		fmt.Printf("recv len:%v msgId:%v err:%v \n", length, responseRegist, err)
	}

	if res.Result == int32(msg.ErrCode_SUCC) {
		for i := 0; i < len(res.Items); i++ {
			fmt.Printf("item:%v\n", res.Items[i])
		}
	}
	fmt.Printf("responseLoadItem result:%v\n", res.Result)

}

func responseUpdateItem(data []byte, length int) {
	res := &msg.NotifyUpdateItem{}
	if err := proto.Unmarshal(data[6:length], res); err != nil {
		fmt.Printf("recv len:%v msgId:%v err:%v\n", length, responseRegist, err)
	}

	fmt.Printf("responseUpdateItem %v\n", res.Items)
}
