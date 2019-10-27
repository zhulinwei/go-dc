package database

import (
	"github.com/zhulinwei/go-dc/pkg/config"
	"github.com/zhulinwei/go-dc/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"sync"
)

var mongoOnce sync.Once
var mongoMutex sync.Mutex
var mongoClientMap map[string]*mongo.Client

type IMongoDB interface {
	UserCollection() *mongo.Collection
}

type MongoDB struct {
	DB1Client *mongo.Client
}

func BuildMongoDB() IMongoDB {
	initMongoDB()
	return MongoDB{
		DB1Client: mongoClientMap["db1"],
	}
}

func (mongodb MongoDB) UserCollection() *mongo.Collection {
	return mongodb.DB1Client.Database("test_database").Collection("test_collection")
}

func initMongoDB() {
	mongoConfigs := config.ServerConfig().MongoDB
	mongoOnce.Do(func() {
		mongoMutex.Lock()
		defer mongoMutex.Unlock()

		mongoClientMap = make(map[string]*mongo.Client, len(mongoConfigs))
		for _, mongoConfig := range mongoConfigs {
			var err error
			var client *mongo.Client

			// 解析mongo链接地址
			content := util.GetHelper().GetContent()
			mongoOptions := options.Client().ApplyURI(mongoConfig.Addr)
			// 连接mongodb数据库
			if client, err = mongo.Connect(content, mongoOptions); err != nil {
				log.Fatalf("mongodb connect failed: %v", err)
				return
			}
			// 检查MongoDB状态值
			if err = client.Ping(content, readpref.Primary()); err != nil {
				log.Fatalf("mongodb ping failed: %v", err)
				return
			}
			// 保存mongodb客户端
			mongoClientMap[mongoConfig.Name] = client
		}
	})
}
