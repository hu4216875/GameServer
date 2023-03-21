package main

import (
	"client/game"
	"client/msg"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
)

var g_userId = "test"

// DisptchMsg
func DisptchMsg(conn net.Conn) {
	recv := make([]byte, 1024*1024)
	var writeIndex = 0
	for {
		len, err := conn.Read(recv[writeIndex:])
		if err == nil {
			if len < 6 {
				continue
			}
			dataLen := binary.BigEndian.Uint32(recv)
			msgId := binary.BigEndian.Uint16(recv[4:])

			// 是一个完整包
			msgLen := int(dataLen + 4)
			if len >= msgLen {
				game.DisptchResponseMsg(msg.MsgId(msgId), recv, len)
				// 多余一个数据包
				if len > msgLen {
					recv = recv[len:]
					writeIndex = len - msgLen
				}
			}
		} else if err != nil {
			if len == 0 {
				fmt.Printf("##################### userId:%v is knick out\n", g_userId)
				conn.Close()
				break
			}
			fmt.Printf("DisptchMsg err %v", err)
		}
	}
}

func NewRequestRegistMsg() *msg.RequestRegist {
	return &msg.RequestRegist{
		UserId: "test",
		Passwd: "123456",
	}
}

func NewRequestLoginMsg() *msg.RequestLogin {
	return &msg.RequestLogin{
		UserId: g_userId,
		Passwd: "123456",
	}
}

func main() {
	game.InitResponseHandler()
	conn, err := net.Dial("tcp", "192.168.5.8:3563")
	if err != nil {
		panic(err)
	}

	msgTemp := NewRequestLoginMsg()

	// 进行编码
	data, err := proto.Marshal(msgTemp)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	// len + data
	m := make([]byte, 2+4+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint32(m, uint32(2+len(data)))

	//binary.BigEndian.PutUint16(m[4:], uint16(msg.MsgId_ID_RequestRegist))
	binary.BigEndian.PutUint16(m[4:], uint16(msg.MsgId_ID_RequestLogin))

	copy(m[6:], data)

	fmt.Println(m)
	// 发送消息
	conn.Write(m)

	fmt.Println("-------send msg-------")

	go DisptchMsg(conn)
	for {
	}
}
