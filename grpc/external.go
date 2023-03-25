package grpc

import (
	"server/grpc/internal"
	"server/grpc/internal/service"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func GetOreEndTime() uint32 {
	return service.GetEndTime()
}
