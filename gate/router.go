package gate

import (
	"server/game"
	"server/login"
	"server/msg"
)

func init() {
	msg.Processor.SetRouter(&msg.RequestRegist{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestLogin{}, login.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestLogout{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestLoadItem{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestGMCommand{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestClientHeart{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestOreTotal{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestStartOre{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestEndOre{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestUpgradeOreSpeed{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestOreInfo{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestEnterBattle{}, game.ChanRPC)
	msg.Processor.SetRouter(&msg.RequestLeaveBattle{}, game.ChanRPC)
}
