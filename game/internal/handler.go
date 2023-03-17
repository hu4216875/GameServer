package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"reflect"
	"server/msg"
)

func init() {
	skeleton.RegisterChanRPC("Register", rpcRegister)
	skeleton.RegisterChanRPC("Login", rpcLogin)
	skeleton.RegisterChanRPC("Logout", rpcLogout)

	handler(&msg.Hello{}, handleHello)
}

func rpcRegister(args []interface{}) {

}

// rpcLogin 登录
func rpcLogin(args []interface{}) {
	//loginMsg := args[0].(*msg.RequestLogin)
	//agent := args[1].(gate.Agent)

	//time.Sleep(120 * time.Second)
	//agent.WriteMsg(&msg.ResponseLogin{
	//	Result: 1,
	//})

	//log.Debug("rpcLogin username:%v", loginMsg.Username)
}

// rpcLogout 退出
func rpcLogout(args []interface{}) {
	username := args[0].(string)
	log.Debug("rpcLogout username:%v", username)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*msg.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %v", m.GetName())
	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&msg.Hello{
		Name: m.GetName(),
	})
}
