package dao

import (
	"errors"
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func (database *Redis) InitRedis() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(errors.New("redis connect fail"))
	}
	database.client = client
}