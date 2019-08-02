package database

import (
	"fmt"
	"github.com/go-redis/redis"
)
import "go.mongodb.org/mongo-driver/mongo"

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
	//database.mongodb = mongo.
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))
	fmt.Println(client, err)
}

func (database *Database) InitDatabase () {
	database.InitRedis()
	database.InitMongoDB()
}