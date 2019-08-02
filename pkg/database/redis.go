package database

import (
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	redis *redis.Client
	mongodb *mongo.Client
}

func (database *Database) InitRedis() {
	database.redis = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func (database *Database) InitMongoDB () {
	//database.mongodb = mongo
}