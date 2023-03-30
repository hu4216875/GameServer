package dao

import (
	"context"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"server/db"
	"server/game/internal/model"
	"server/publicconst"
)

var (
	ServerInfoDao = &serverInfoDao{}
)

type serverInfoDao struct {
}

func (s *serverInfoDao) ExistServerInfo(serverId uint32) bool {
	collection := db.GetGlobalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.GLOBAL_SERVERINFO_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	result := collection.FindOne(ctx, bson.M{"serverid": serverId})
	if result == nil {
		return false
	}

	var serverInfo model.ServerInfo
	if err := result.Decode(&serverInfo); err != nil {
		return false
	}
	return true
}

// AddServerInfo 添加服务器
func (s *serverInfoDao) AddServerInfo(serverId uint32) {
	serverInfo := model.NewServerInfo(serverId)
	collection := db.GetGlobalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.GLOBAL_SERVERINFO_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.InsertOne(ctx, serverInfo); err != nil {
		log.Error("AddServerInfo err:%v", err)
	}
}

// UpdateRegistNum 更新注册数
func (s *serverInfoDao) UpdateRegistNum(serverId uint32) {
	collection := db.GetGlobalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.GLOBAL_SERVERINFO_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.UpdateOne(ctx, bson.M{"serverid": serverId}, bson.D{{"$inc", bson.D{{"registnum", 1}}}}); err != nil {
		log.Error("UpdateServerTime err:%v", err)
	}
}
