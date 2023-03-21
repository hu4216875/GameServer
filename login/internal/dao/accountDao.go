package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"server/db"
	"server/login/internal/model"
	"server/msg"
	"server/publicconst"
)

var (
	AccountDao = NewAcccountDao()
)

type acccountDao struct {
}

func NewAcccountDao() *acccountDao {
	return &acccountDao{}
}

func (a *acccountDao) Regist(userId, pwd string, accountId int64) (msg.ErrCode, *model.Account) {
	account := model.NewAccount(userId, pwd, accountId)
	collection := db.GetGlobalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.GLOBAL_ACCOUNT_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.InsertOne(ctx, account); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return msg.ErrCode_USER_ID_EXIST, nil
		} else {
			return msg.ErrCode_SYSTEM_ERROR, nil
		}
	}
	return msg.ErrCode_SUCC, account
}

// FindAccount 账户是否存在
func (a *acccountDao) FindAccount(userId string) *model.Account {
	collection := db.GetGlobalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.GLOBAL_ACCOUNT_COLLECTION)

	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
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
