package database

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type Databse struct {}

type RedisConfig struct {
	Url string
}

type MongoConfig struct {
	Url            string
	DatabaseName   string
	CollectionName string
}

var cache = new(Redis)
var mongodb = new(MongoDB)

func GetReis() *redis.Client {
	return cache.client
}

func GetMongoDBCursor() *mongo.Cursor  {
	return mongodb.cursor
}

func GetMongoDBClient() *mongo.Client {
	return mongodb.client
}

func GetMongoDBCollection () *mongo.Collection {
	return mongodb.collection
}

func init() {
	// TODO 配置应该动态读取
	cache.InitRedis(&RedisConfig{
		Url: "localhost:6379",
	})
	mongodb.InitMongoDB(&MongoConfig{
		Url:            "mongodb://localhost:27017",
		DatabaseName:   "test",
		CollectionName: "test",
	})
}