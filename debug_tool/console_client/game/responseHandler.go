package game

import (
	"client/msg"
	"fmt"
	"net"
)

var g_conn net.Conn

type ResponseFunc func(data []byte, len int)

var responseHandler map[msg.MsgId]ResponseFunc

func InitResponseHandler(conn net.Conn) {
	g_conn = conn
	responseHandler = make(map[msg.MsgId]ResponseFunc)
	registAccountResponse()
	registItemResponse()
	registCmCommandResponse()
	registOreResponse()
}

func DisptchResponseMsg(msgId msg.MsgId, data []byte, len int) {
	fmt.Printf("DisptchResponseMsg msgId:%v \n", msgId)
	if fn, ok := responseHandler[msgId]; ok {
		fn(data, len)
	} else {
		fmt.Printf("DisptchResponseMsg msgId:%v not found\n", msgId)
	}
}

func registResponse(msgId msg.MsgId, fn ResponseFunc) {
	responseHandler[msgId] = fn
}
