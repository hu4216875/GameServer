package dao

import (
	"context"
	"github.com/name5566/leaf/log"
	"server/db"
	"server/game/internal/model"
	"server/publicconst"
)

var (
	ItemLogDao = &itemLogDao{}
)

type itemLogDao struct {
}

func (i *itemLogDao) AddItemLog(itemLog *model.ItemLog) {
	collection := db.GetLocalClient().Database(publicconst.GLOBAL_DB_NAME).Collection(publicconst.LOG_DB_NAME)
	ctx, cancel := context.WithTimeout(context.Background(), publicconst.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.InsertOne(ctx, itemLog); err != nil {
		log.Error("AddItemLog err:%v", err)
	}
}
