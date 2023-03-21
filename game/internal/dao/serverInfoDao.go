package dao

import (
	"context"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"server/db"
	"server/game/internal/model"
	"server/publicconst"
	"time"
)

var (
	ServerInfoDao = &serverInfoDao{}
)

type serverInfoDao struct {
}

func (s *serverInfoDao) ExistServerInfo(serverAddr string) bool {
	collection := db.GetGlobalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.GLOBAL_SERVERINFO_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	result := collection.FindOne(ctx, bson.M{"serveraddr": serverAddr})
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
func (s *serverInfoDao) AddServerInfo(serverAddr string) {
	serverInfo := model.NewServerInfo(serverAddr)
	collection := db.GetGlobalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.GLOBAL_SERVERINFO_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.InsertOne(ctx, serverInfo); err != nil {
		log.Error("AddServerInfo err:%v", err)
	}
}

// UpdateServerTime 更新服务器时间
func (s *serverInfoDao) UpdateServerTime(serverAddr string) {
	curTime := uint32(time.Now().Unix())
	collection := db.GetGlobalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.GLOBAL_SERVERINFO_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.UpdateOne(ctx, bson.M{"serveraddr": serverAddr}, bson.D{{"$set", bson.D{{"updatetime", curTime}}}}); err != nil {
		log.Error("UpdateServerTime err:%v", err)
	}
}

// UdpateRegistNum 更新注册数
func (s *serverInfoDao) UdpateRegistNum(serverAddr string) {
	collection := db.GetGlobalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.GLOBAL_SERVERINFO_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.UpdateOne(ctx, bson.M{"serveraddr": serverAddr}, bson.D{{"$inc", bson.D{{"registnum", 1}}}}); err != nil {
		log.Error("UpdateServerTime err:%v", err)
	}
}
