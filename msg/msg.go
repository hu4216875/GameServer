package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(&RequestGMCommand{})
	Processor.Register(&ResponseGMCommand{})
	Processor.Register(&RequestRegist{})
	Processor.Register(&ResponseRegist{})
	Processor.Register(&RequestLogin{})
	Processor.Register(&ResponseLogin{})
	Processor.Register(&RequestLogout{})
	Processor.Register(&ResponseLogout{})
	Processor.Register(&ResponseKickOut{})
	Processor.Register(&RequestLoadItem{})
	Processor.Register(&ResponseLoadItem{})
	Processor.Register(&NotifyUpdateItem{})
	Processor.Register(&RequestClientHeart{})
	Processor.Register(&ResponseClientHert{})
	Processor.Register(&RequestOreTotal{})
	Processor.Register(&ResponseOreTotal{})
	Processor.Register(&RequestStartOre{})
	Processor.Register(&ResponseStartOre{})
	Processor.Register(&RequestEndOre{})
	Processor.Register(&ResponseEndOre{})
	Processor.Register(&RequestUpgradeOreSpeed{})
	Processor.Register(&ResponseUpgradeOreSpeed{})
	Processor.Register(&RequestOreInfo{})
	Processor.Register(&ResponseOreInfo{})
}
