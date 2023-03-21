package game

import (
	"client/msg"
	"fmt"
)

type ResponseFunc func(data []byte, len int)

var responseHandler map[msg.MsgId]ResponseFunc

func InitResponseHandler() {
	responseHandler = make(map[msg.MsgId]ResponseFunc)
	registAccountResponse()
}

func DisptchResponseMsg(msgId msg.MsgId, data []byte, len int) {
	if fn, ok := responseHandler[msgId]; ok {
		fn(data, len)
	} else {
		fmt.Println("DisptchResponseMsg msgId:%d not found", msgId)
	}
}

func registResponse(msgId msg.MsgId, fn ResponseFunc) {
	responseHandler[msgId] = fn
}
