package msg

import (
	"github.com/name5566/leaf/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	//Processor.Register(&RequestLogin{})
	//Processor.Register(&ResponseLogin{})
	Processor.Register(&Hello{})
}
