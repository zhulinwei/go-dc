package database

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

type MongoDB struct {
	client *mongo.Client
	cursor *mongo.Cursor
	collection *mongo.Collection
}

func getContext() (ctx context.Context) {
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	return ctx
}


func (database *MongoDB) InitMongoDB(mongoConfig *MongoConfig) {
	// TODO 需要检查传入的配置
	var err error
	var client *mongo.Client

	// 设置MongoDB选项值
	mongoOptions := options.Client().ApplyURI(mongoConfig.Url)
	// 链接MongoDB数据库
	if client, err = mongo.Connect(getContext(), mongoOptions); err != nil {
		panic(errors.New("mongodb connect fail"))
	}
	// 检查MongoDB状态值
	if err = client.Ping(getContext(), readpref.Primary()); err != nil {
		panic(errors.New("mongodb ping fail"))
	}

	database.client = client
	database.collection = client.Database(mongoConfig.DatabaseName).Collection(mongoConfig.CollectionName)
}
