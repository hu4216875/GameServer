package dao

import (
	"errors"
	"fmt"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"oreDistrictServer/common"
	"oreDistrictServer/db"
	"oreDistrictServer/model"
	"time"
)

var (
	OreDistrictDao oreDistrictDao
)

type oreDistrictDao struct {
}

// LoadOreDistrict 加载所有矿洞
func (o *oreDistrictDao) LoadOreDistrict() []*model.OreDistrict {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.Ore_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	cur, currErr := collection.Find(ctx, bson.D{})
	if currErr != nil {
		log.Error("LoadOreDistrict err:%v", currErr)
		return nil
	}
	defer cur.Close(ctx)

	var oreDistrict []*model.OreDistrict
	if err := cur.All(ctx, &oreDistrict); err != nil {
		log.Error("LoadOreDistrict err:%v", err)
		return nil
	}
	return oreDistrict
}

// AddOreDistrict 添加矿洞
func (o *oreDistrictDao) AddOreDistrict(data *model.OreDistrict) error {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.Ore_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	if _, err := collection.InsertOne(ctx, data); err != nil {
		return errors.New(fmt.Sprintf("AddOreDistrict err:%v", err))
	}

	oreLog := model.NewOreDistrictLog(data.OreDistId)
	logCollection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.Ore_DB_LOG_COLLECTION)
	opts := options.UpdateOptions{}
	opts.SetUpsert(true)
	if _, err := logCollection.UpdateOne(ctx, bson.M{"oreid": data.OreDistId}, bson.D{{"$set", oreLog}}, &opts); err != nil {
		return errors.New(fmt.Sprintf("AddOreDistrict err:%v", err))
	}
	return nil
}

// AddOreDistrictPlayer 添加矿洞玩家
func (o *oreDistrictDao) AddOreDistrictPlayer(ore *model.OreDistrict, player *model.OreDistrictPlayer) error {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.Ore_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	curTime := uint32(time.Now().Unix())
	filter := bson.M{"oredistid": ore.OreDistId}
	update := bson.D{{"$addToSet", bson.D{{"players", player}}},
		{"$set", bson.D{{"endtime", ore.EndTime}, {"updatetime", curTime}}}}
	if _, err := collection.UpdateOne(ctx, filter, update); err != nil {
		log.Error("AddOreDistrictPlayer err:%v", err)
		return errors.New(fmt.Sprintf("AddOreDistrictPlayer err:%v", err))
	}
	return nil
}

// RemoveOreDistrictPlayer 移除矿洞玩家
func (o *oreDistrictDao) RemoveOreDistrictPlayer(ore *model.OreDistrict, player *model.OreDistrictPlayer) error {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.Ore_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	curTime := uint32(time.Now().Unix())
	update := bson.D{{"$set", bson.D{{"endtime", ore.EndTime}, {"total", ore.Total}, {"updatetime", curTime}}}}
	if _, err := collection.UpdateOne(ctx, bson.D{{"oredistid", ore.OreDistId}}, update); err != nil {
		log.Error("Account:%v UpdateOreDistrictPlayer error:%v", player.AccountId, err)
		return err
	}

	if _, err := collection.UpdateOne(ctx, bson.M{"oredistid": ore.OreDistId}, bson.M{"$pull": bson.M{"players": bson.M{"accountid": player.AccountId}}}); err != nil {
		log.Error("Account:%v RemoveOreDistrictPlayer err:%v", player.AccountId, err)
		return err
	}
	return nil
}

// UpdateOreDistrictPlayer 更新矿洞玩家
func (o *oreDistrictDao) UpdateOreDistrictPlayer(ore *model.OreDistrict, player *model.OreDistrictPlayer) error {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.Ore_DB_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()

	curTime := uint32(time.Now().Unix())
	update := bson.D{{"$set", bson.D{{"endtime", ore.EndTime}, {"total", ore.Total}, {"updatetime", curTime}}}}

	if _, err := collection.UpdateOne(ctx, bson.D{{"oredistid", ore.OreDistId}}, update); err != nil {
		log.Error("Account:%v UpdateOreDistrictPlayer error:%v", player.AccountId, err)
		return err
	}

	filter := bson.M{"oredistid": ore.OreDistId, "players.accountid": player.AccountId}
	update = bson.D{{"players.$[item].speed", player.Speed}, {"players.$[item].starttime", player.StartTime}}
	arrayFilter := bson.M{"item.accountid": player.AccountId}
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
		log.Error("Account:%v UpdateOreDistrictPlayer error:%v", player.AccountId, res.Err())
	}
	return nil
}

// UpdateOreRecord 更新矿洞记录
func (o *oreDistrictDao) UpdateOreRecord(oreId uint32, accountId int64, num uint32) {
	collection := db.GlobalClient.Database(common.ORE_DB_NAME).Collection(common.Ore_DB_LOG_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), common.DB_OP_TIME_OUT)
	defer cancel()
	record := model.NewOreDistrictRecord(accountId, num)
	if _, err := collection.UpdateOne(ctx, bson.M{"oreid": oreId}, bson.D{{"$addToSet", bson.D{{"records", record}}}}); err != nil {
		log.Error("UpdateOreRecord err:%v", err)
	}
}
