package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"server/db"
	"server/game/internal/model"
	"server/msg"
	"server/publicconst"
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
