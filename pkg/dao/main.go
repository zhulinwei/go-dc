package dao

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type Dao struct {
	Redis
	MongoDB
}

var dao = new(Dao)

func GetReis() *redis.Client {
	return dao.Redis.client
}

func GetMongoDB() *mongo.Client {
	return dao.MongoDB.client
}

func init() {
	dao.Redis.InitRedis()
	dao.MongoDB.InitMongoDB()
}
