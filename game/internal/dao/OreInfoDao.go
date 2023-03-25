package dao

import (
	"context"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"server/db"
	"server/game/internal/model"
	"server/msg"
	"server/publicconst"
)

var (
	OreInfoDao = &oreInfoDao{}
)

type oreInfoDao struct {
}

func (o *oreInfoDao) UpdateOreInfo(accountId int64, oreInfo *model.OreInfo) msg.ErrCode {
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOCAL_ACCOUNT_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()
	if _, err := collection.UpdateOne(ctx, bson.M{"accountid": accountId},
		bson.D{{"$set", bson.D{{"oreinfo.oreid", oreInfo.OreId}}},
			{"$set", bson.D{{"oreinfo.starttime", oreInfo.StartTime}}},
			{"$set", bson.D{{"oreinfo.speed", oreInfo.Speed}}}}); err != nil {
		log.Error("AddItem err:%v", err)
		return msg.ErrCode_SYSTEM_ERROR
	}
	return msg.ErrCode_SUCC
}
