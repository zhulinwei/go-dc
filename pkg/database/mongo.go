package database

import (
	"sync"

	"github.com/zhulinwei/go-dc/pkg/config"

	"github.com/zhulinwei/go-dc/pkg/util"
	"github.com/zhulinwei/go-dc/pkg/util/log"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/zhulinwei/go-dc/pkg/model"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	databaseKey    = "db1"
	databaseName   = "test_database"
	collectionName = "test_collection"
)

type IMongoDB interface {
	UserCollection() *mongo.Collection
}

type MongoDB struct {
	once      sync.Once
	configs   []model.MongoConfig
	clientMap map[string]*mongo.Client
}

var mongodb *MongoDB

func BuildMongoDB() IMongoDB {
	if mongodb == nil {
		mongodb = &MongoDB{configs: config.ServerConfig().MongoDB}
		mongodb.init()
	}
	return mongodb
}

func (mongodb *MongoDB) UserCollection() *mongo.Collection {
	return mongodb.clientMap[databaseKey].Database(databaseName).Collection(collectionName)
}

func (mongodb *MongoDB) init() {
	mongodb.once.Do(func() {
		mongodb.clientMap = make(map[string]*mongo.Client, len(mongodb.configs))
		for _, mongoConfig := range mongodb.configs {
			clientOptions := options.Client().ApplyURI(mongoConfig.Addr)
			mongoClient, err := mongo.Connect(util.CommonContent(), clientOptions)
			if err != nil {
				log.Error("mongodb connection fail", log.String("error", err.Error()))
				panic("mongodb connection fail")
			}
			mongodb.clientMap[mongoConfig.Name] = mongoClient
		}
	})
}
