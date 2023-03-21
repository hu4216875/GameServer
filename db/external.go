package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"server/db/internal"
)

var (
	Module  = new(internal.Module)
	ChanRPC = internal.ChanRPC
)

func GetGlobalClient() *mongo.Client {
	return internal.GlobalClient
}

func GetLocalClient() *mongo.Client {
	return internal.LocalClient
}
