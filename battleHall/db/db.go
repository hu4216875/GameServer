package db

import (
	"battleHall/conf"
	"github.com/name5566/leaf/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"time"
)

var GlobalClient *mongo.Client

func OnInit() {
	GlobalClient = initDB(conf.Server.MongoDbUrl, conf.Server.MaxMongoDbConn)
}

func OnRelease() {
	if GlobalClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		GlobalClient.Disconnect(ctx)
	}
}

func initDB(dburl string, maxConn uint32) *mongo.Client {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(dburl)
	clientOptions.SetMaxConnecting(uint64(maxConn))
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("initDB err：%v", err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("initDB client ping err：%v", err)
	}
	return client
}
