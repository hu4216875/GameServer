package game

import (
	"client/msg"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
)

func SendRegistMsg(userId string) {
	msgTemp := &msg.RequestRegist{
		UserId: userId,
		Passwd: "123456",
	}

	// 进行编码
	data, err := proto.Marshal(msgTemp)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestRegist), data)
}

func SendLoginMsg(userId string) {
	msgTemp := &msg.RequestLogin{
		UserId: userId,
		Passwd: "123456",
	}

	// 进行编码
	data, err := proto.Marshal(msgTemp)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestLogin), data)
}

func SendLogoutMsg() {
	msgTemp := &msg.RequestLogout{}

	// 进行编码
	data, err := proto.Marshal(msgTemp)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestLogout), data)
}

func SendRequestLoadItem(conn net.Conn) {
	loadItem := &msg.RequestLoadItem{}
	// 进行编码
	data, err := proto.Marshal(loadItem)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestLoadItem), data)
}

func sendClientHert() {
	hertClientMsg := &msg.RequestClientHeart{}
	// 进行编码
	data, err := proto.Marshal(hertClientMsg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestClientHeart), data)
}

func sendMsg(msgId uint16, data []byte) {
	// len + data
	m := make([]byte, 2+4+len(data))
	// 默认使用大端序
	binary.BigEndian.PutUint32(m, uint32(2+len(data)))
	binary.BigEndian.PutUint16(m[4:], msgId)
	copy(m[6:], data)
	// 发送消息
	g_conn.Write(m)
}

func SendRequestOreTotal() {
	hertClientMsg := &msg.RequestOreTotal{}
	// 进行编码
	data, err := proto.Marshal(hertClientMsg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestOreTotal), data)
}

func SendStartOre() {
	startOreMsg := &msg.RequestStartOre{}
	// 进行编码
	data, err := proto.Marshal(startOreMsg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestStartOre), data)
}

func SendEndOre() {
	startOreMsg := &msg.RequestEndOre{}
	// 进行编码
	data, err := proto.Marshal(startOreMsg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestEndOre), data)
}

func SendChangeSpeed() {
	changeSpeed := &msg.RequestUpgradeOreSpeed{}
	data, err := proto.Marshal(changeSpeed)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestUpgradeOreSpeed), data)
}

func SendEnterBattle() {
	enterBattle := &msg.RequestEnterBattle{}
	data, err := proto.Marshal(enterBattle)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	sendMsg(uint16(msg.MsgId_ID_RequestEnterBattle), data)
}
