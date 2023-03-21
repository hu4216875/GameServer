package internal

import (
	"reflect"
)

type HandleFunc func(args []interface{})

func init() {
	skeleton.RegisterChanRPC("Regist", rpcRegist)
	skeleton.RegisterChanRPC("Login", rpcLogin)
	skeleton.RegisterChanRPC("Logout", rpcLogout)

}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func makeHandleMsg(f HandleFunc) func(args []interface{}) {
	return func(args []interface{}) {
		f(args)
	}
}

func handleHello(args []interface{}) {
	//// 收到的 Hello 消息
	//m := args[0].(*msg.Hello)
	//// 消息的发送者
	//a := args[1].(gate.Agent)
	//
	//// 输出收到的消息的内容
	//log.Debug("hello %v", m.GetName())
	//// 给发送者回应一个 Hello 消息
	//a.WriteMsg(&msg.Hello{
	//	Name: m.GetName(),
	//})
}
