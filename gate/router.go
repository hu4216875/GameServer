package gate

import (
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.RequestRegist{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestLogin{}, login.ChanRPC)
}
