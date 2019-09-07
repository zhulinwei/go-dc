package database

import (
	"errors"
	"github.com/zhulinwei/gin-demo/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type IMongoDB interface {
	UserCollection() *mongo.Collection
}

type MongoDB struct {
	client *mongo.Client
}

func (mongodb MongoDB) UserCollection () *mongo.Collection {
	return mongodb.client.Database("test_database").Collection("test_collection")
}

func BuildMongoDB () IMongoDB {
	return MongoDB{
		client: connectMongoDB(&MongoConfig{
			Url:            "mongodb://localhost:27017",
			DatabaseName:   "test",
			CollectionName: "test",
		}),
	}
}

func connectMongoDB(mongoConfig *MongoConfig) *mongo.Client {
	var err error
	var client *mongo.Client

	// 链接MongoDB数据库
	content := util.GetHelper().GetContent()
	// 设置MongoDB选项值
	mongoOptions := options.Client().ApplyURI(mongoConfig.Url)
	if client, err = mongo.Connect(content, mongoOptions); err != nil {
		panic(errors.New("mongodb connect fail"))
	}
	// 检查MongoDB状态值
	if err = client.Ping(content, readpref.Primary()); err != nil {
		panic(errors.New("mongodb ping fail"))
	}
	return client
}