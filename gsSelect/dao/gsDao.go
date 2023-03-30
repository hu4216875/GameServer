package dao

import (
	"context"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"gsSelect/common"
	"gsSelect/db"
	"gsSelect/model"
)

var (
	GsDao = &gsDao{}
)

type gsDao struct {
}

func (g *gsDao) LoadAllGs() []*model.ServerInfo {
	collection := db.GlobalClient.Database(common.GLOBAL_DB_NAME).Collection(common.GLOBAL_SERVERINFO_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	cur, currErr := collection.Find(ctx, bson.D{})
	if currErr != nil {
		log.Error("LoadAllGs currErr:%v", currErr)
	}
	defer cur.Close(ctx)
	var ret []*model.ServerInfo
	if err := cur.All(ctx, &ret); err != nil {
		log.Error("LoadAllGs err:%v", err)
	}
	return ret
}

func (g *gsDao) FindAccount(userId string) *model.Account {
	collection := db.GlobalClient.Database(common.GLOBAL_DB_NAME).Collection(common.GLOBAL_ACCOUNT_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	result := collection.FindOne(ctx, bson.M{"userid": userId})
	if result == nil {
		return nil
	}

	var account model.Account
	if err := result.Decode(&account); err != nil {
		return nil
	}
	return &account
}
