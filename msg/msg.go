package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
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
}
