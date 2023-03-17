package main

import (
	"encoding/binary"
	"net"
	"fmt"
	"github.com/golang/protobuf/proto"
	"client/msg"
	"log"
)

// msgId 2 dataLen 4 data
func DisptchMsg(conn net.Conn) {
	recv := make([]byte, 30)
	for {
		if len, err := conn.Read(recv);err==nil {
			if len < 6 {
				continue
			}
			dataLen := binary.BigEndian.Uint32(recv)
			msgId := binary.BigEndian.Uint16(recv[4:])

			// 是一个完整包
			if int(dataLen + 4) == len {
				temp := &msg.Hello{}
			    if err := proto.Unmarshal(recv[6:len], temp);err == nil {
			        fmt.Printf("recv len:%v msgId:%v, content:%s", len,  msgId, temp.Name)
			    }
			}
		}
	}
}


func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	temp := &msg.Hello {
		Name : "test",
	}

	  // 进行编码
    data, err := proto.Marshal(temp)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }

	// len + data
	m := make([]byte, 2 + 4 + len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint32(m, uint32(2 + len(data)))
	binary.BigEndian.PutUint16(m, 0)

	copy(m[6:], data)

	// 发送消息
	conn.Write(m)

	go DisptchMsg(conn)

	for{
	}
}


