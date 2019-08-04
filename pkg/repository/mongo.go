package repository

import (
	"errors"
	. "github.com/zhulinwei/gin-demo/pkg/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	client     *mongo.Client
	cursor     *mongo.Cursor
	collection *mongo.Collection
}

func (database *MongoDB) InitMongoDB(mongoConfig *MongoConfig) {
	// TODO 需要检查传入的配置
	var err error
	var client *mongo.Client

	// 链接MongoDB数据库
	content := GetUitl().GetContent()
	// 设置MongoDB选项值
	mongoOptions := options.Client().ApplyURI(mongoConfig.Url)
	if client, err = mongo.Connect(content, mongoOptions); err != nil {
		panic(errors.New("mongodb connect fail"))
	}
	// 检查MongoDB状态值
	if err = client.Ping(content, readpref.Primary()); err != nil {
		panic(errors.New("mongodb ping fail"))
	}

	database.client = client
	database.collection = client.Database(mongoConfig.DatabaseName).Collection(mongoConfig.CollectionName)
}
