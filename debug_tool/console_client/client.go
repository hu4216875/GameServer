package main

import (
	"client/game"
	"client/msg"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

var g_userId = "test"
var g_conn net.Conn

// DisptchMsg
func DisptchMsg(conn net.Conn, ch chan<- int) {
	recv := make([]byte, 1024*1024)
	var writeIndex = 0
	for {
		readLen, err := conn.Read(recv[writeIndex:])
		//	fmt.Printf("########### readLen:%v\n", readLen)
		if err == nil {
			for {
				if readLen < 6 {
					break
				}
				dataLen := binary.BigEndian.Uint32(recv)
				msgId := binary.BigEndian.Uint16(recv[4:])

				//	fmt.Printf("------msgId:%v readLen:%v\n", msgId, readLen)
				// 是一个完整包
				msgLen := int(dataLen + 4)
				if readLen >= msgLen {
					game.DisptchResponseMsg(msg.MsgId(msgId), recv, readLen)
					// 多余一个数据包
					if readLen > msgLen {
						recv = recv[msgLen:]
						writeIndex = readLen - msgLen
					} else {
						writeIndex = 0
					}
					readLen = readLen - msgLen
				} else {
					break
				}
			}

		} else if err != nil {
			if readLen == 0 {
				fmt.Printf("##################### userId:%v is knick out\n", g_userId)
				conn.Close()
				ch <- 1
				break
			}
			fmt.Printf("DisptchMsg err %v", err)
		}
	}
}

func main() {
	var err error
	g_conn, err = net.Dial("tcp", "192.168.5.8:3563")
	if err != nil {
		panic(err)
	}
	game.InitResponseHandler(g_conn)

	//game.SendRegistMsg(g_userId)
	game.SendLoginMsg(g_userId)
	time.Sleep(1000)
	game.SendRequestLoadItem(g_conn)
	time.Sleep(1000)

	game.SendEnterBattle()
	//game.SendGmReload()

	//game.SendLogoutMsg()
	//	game.SendStartOre()

	//game.SendStartOre()
	//game.SendChangeSpeed()
	//game.SendEndOre()
	//game.SendGmAddItem()
	//	game.SendRequestOreTotal()

	ch := make(chan int, 1)
	go DisptchMsg(g_conn, ch)
	<-ch
	fmt.Println("exit")
}
