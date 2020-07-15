package database

import (
	"fmt"
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
		var wg sync.WaitGroup

		length := len(mongodb.configs)
		wg.Add(length)
		mongodb.clientMap = make(map[string]*mongo.Client, length)

		for _, mongoConfig := range mongodb.configs {

			go func(config model.MongoConfig, wg *sync.WaitGroup) {
				defer wg.Done()

				clientOptions := options.Client().ApplyURI(config.Addr)
				// clientOptions.SetAuth(options.Credential{Username: "admin", Password: "admin"})
				fmt.Println(clientOptions.Auth)
				mongoClient, err := mongo.Connect(util.CommonContent(), clientOptions)
				if err != nil {
					log.Error("mongodb connection fail", log.String("error", err.Error()))
					return
				}

				mongodb.clientMap[config.Name] = mongoClient
			}(mongoConfig, &wg)
		}

		wg.Wait()
	})
}
