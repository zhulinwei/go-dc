package database

import (
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func (database *Redis) InitRedis() {
	database.client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}