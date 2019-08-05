package database

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type RedisConfig struct {
	Url string
}

type MongoConfig struct {
	Url            string
	DatabaseName   string
	CollectionName string
}

func InitDatabase () (*redis.Client, *mongo.Client) {
	cache := new(Cache)
	cacheInstance := cache.InitRedis(&RedisConfig{
		Url: "localhost:6379",
	})

	mongodb := new(MongoDB)
	mongodbInstance := mongodb.InitMongoDB(&MongoConfig{
		Url:            "mongodb://localhost:27017",
		DatabaseName:   "test",
		CollectionName: "test",
	})

	return cacheInstance, mongodbInstance
}