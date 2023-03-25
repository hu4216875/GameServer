package dao

import (
	"context"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/db"
	"server/game/internal/model"
	"server/msg"
	"server/publicconst"
	"server/util"
)

var (
	AccountDao = &accountDao{}
)

type accountDao struct {
}

func (a *accountDao) AddAccount(userId string, accountId int64) msg.ErrCode {
	account := model.NewAccount(userId, accountId)
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOCAL_ACCOUNT_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.InsertOne(ctx, account); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return msg.ErrCode_USER_ID_EXIST
		} else {
			return msg.ErrCode_SYSTEM_ERROR
		}
	}

	return msg.ErrCode_SUCC
}

func (a *accountDao) GetAccount(accountId int64) *model.Account {
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOCAL_ACCOUNT_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	result := collection.FindOne(ctx, bson.M{"accountid": accountId})
	if result == nil {
		return nil
	}

	var ret model.Account
	err := result.Decode(&ret)
	if err != nil {
		log.Error("Account:%v GetAccount err:%v", accountId, err)
	}
	return &ret
}

func (a *accountDao) UpdateAccountLogin(accountId int64) {
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOCAL_ACCOUNT_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.UpdateOne(ctx, bson.M{"accountid": accountId}, bson.D{{"$set", bson.D{{"logintime", util.GetCurTime()}}}}); err != nil {
		log.Error("UpdateAccountLogin accountId:%v err:%v", accountId, err)
	}
}

func (a *accountDao) UpdateAccountLogout(accountId int64) {
	collection := db.GetLocalClient().Database(publicconst.LOCAL_DB_NAME).Collection(publicconst.LOCAL_ACCOUNT_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.UpdateOne(ctx, bson.M{"accountid": accountId}, bson.D{{"$set", bson.D{{"logouttime", util.GetCurTime()}}}}); err != nil {
		log.Error("UpdateAccountLogin accountId:%v err:%v", accountId, err)
	}
}
