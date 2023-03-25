package dao

import (
	"context"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/db"
	"server/game/internal/model"
	"server/msg"
	"server/publicconst"
	"strings"
	"time"
)

var (
	ItemDao = &itemDao{}
)

type itemDao struct {
}

func (a *itemDao) LoadItem(accountId int64) (msg.ErrCode, *model.AccountItem) {
	collection := db.GetGlobalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOACL_Item)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	result := collection.FindOne(ctx, bson.M{"accountid": accountId})
	if result == nil {
		return msg.ErrCode_SYSTEM_ERROR, nil
	}

	var ret model.AccountItem
	err := result.Decode(&ret)
	if err != nil {
		if strings.Contains(publicconst.MONGO_NO_RESULT, err.Error()) {
			accountItem := model.NewAccountItem(accountId)
			if _, err = collection.InsertOne(ctx, accountItem); err != nil {
				log.Error("LoadItem accountId:%v err:%v", accountId, err)
			}
		} else {
			log.Error("LoadItem accountId:%v err:%v", accountId, err)
			return msg.ErrCode_SYSTEM_ERROR, nil
		}
	}
	return msg.ErrCode_SUCC, &ret
}

func (a *itemDao) AddItem(accountId int64, item *model.Item) msg.ErrCode {
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOACL_Item)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.UpdateOne(ctx, bson.M{"accountid": accountId}, bson.D{{"$addToSet", bson.D{{"items", item}}}}); err != nil {
		log.Error("AddItem err:%v", err)
	}
	return msg.ErrCode_SUCC
}

func (a *itemDao) UpdateItem(accountId int64, item *model.Item) msg.ErrCode {
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOACL_Item)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if item.Num == 0 {
		if _, err := collection.UpdateOne(ctx, bson.M{"accountid": accountId}, bson.M{"$pull": bson.M{"items": bson.M{"id": item.Id}}}); err != nil {
			log.Error("account:%v UpdateItem err:%v", accountId, err)
		} else {
			log.Debug("account:%v delete item:%d", accountId, item)
		}
	} else {
		curTime := uint32(time.Now().Unix())
		filter := bson.M{"accountid": accountId, "items.id": item.Id}
		update := bson.M{"items.$[item].num": item.Num, "items.$[item].limitDate": item.LimitDate, "items.$[item].updatetime": curTime}
		arrayFilter := bson.M{"item.id": item.Id}
		res := collection.FindOneAndUpdate(ctx, filter, bson.M{"$set": update},
			options.FindOneAndUpdate().SetArrayFilters(
				options.ArrayFilters{
					Filters: []interface{}{
						arrayFilter,
					},
				},
			),
		)
		if res.Err() != nil {
			log.Error("account:%v UpdateItem error:%v", accountId, res.Err())
		}
	}
	return msg.ErrCode_SUCC
}
