package dao

import (
	"battleHall/common"
	"battleHall/db"
	"battleHall/model"
	"context"
	"errors"
	"fmt"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ServerInfoDao serverInfoDao
)

type serverInfoDao struct {
}

func (s *serverInfoDao) LoadAll() []*model.ServerInfo {
	collection := db.GlobalClient.Database(common.DB_BATTLE_HALL).Collection(common.DB_BATTLE_HALL_SERVER)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	cur, currErr := collection.Find(ctx, bson.D{})
	if currErr != nil {
		log.Error("LoadOreDistrict err:%v", currErr)
		return nil
	}
	defer cur.Close(ctx)

	var serverInfo []*model.ServerInfo
	if err := cur.All(ctx, &serverInfo); err != nil {
		log.Error("LoadOreDistrict err:%v", err)
		return nil
	}
	return serverInfo
}

func (s *serverInfoDao) InsertOrUpdateServerInfo(data *model.ServerInfo) error {
	collection := db.GlobalClient.Database(common.DB_BATTLE_HALL).Collection(common.DB_BATTLE_HALL_SERVER)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	opts := options.UpdateOptions{}
	opts.SetUpsert(true)
	if _, err := collection.UpdateOne(ctx, bson.M{"serverAddr": data.ServerAddr}, bson.D{{"$set", data}}, &opts); err != nil {
		return errors.New(fmt.Sprintf("AddOreDistrict err:%v", err))
	}
	return nil
}
