package internal

import (
	"reflect"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	//handler(&msg.RequestLogin{}, handleLoginReq)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLoginReq(args []interface{}) {
	//m := args[0].(*msg.RequestLogin)
	//agent := args[1].(gate.Agent)

	//log.Debug("handleLoginReq user:%v login", m.Username)

	// 开启携程第三方验证
	skeleton.Go(func() {
		// 输出收到的消息的内容
		//	log.Debug("handleLoginReq user:%v", m.Username)

		//	game.ChanRPC.Call0("Login", m, agent)
	}, func() {
		//	log.Debug("handleLoginReq user:%v succ", m.Username)
	})
}
