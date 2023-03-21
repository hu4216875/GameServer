package internal

import (
	"context"
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/module"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"server/base"
	"server/conf"
	"time"
)

var (
	skeleton     = base.NewSkeleton()
	ChanRPC      = skeleton.ChanRPCServer
	LocalClient  *mongo.Client
	GlobalClient *mongo.Client
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
	LocalClient = m.initDB(conf.Server.LocalMongoDBUrl, conf.Server.MaxLocalMongoDBConn)
	GlobalClient = m.initDB(conf.Server.GlobalMongoDBUrl, conf.Server.MaxGlobalMongoDBConn)
}

func (m *Module) OnDestroy() {
	if LocalClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		LocalClient.Disconnect(ctx)
	}
	if GlobalClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		GlobalClient.Disconnect(ctx)
	}
}

func (m *Module) initDB(dburl string, maxConn int) *mongo.Client {
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
