package database

import (
	"errors"
	"github.com/go-redis/redis"
)

type Redis struct {
	client *redis.Client
}

func (database *Redis) InitRedis(redisConfig *RedisConfig) {
	var redisOptions = &redis.Options{
		Addr: redisConfig.Url,
	}
	client := redis.NewClient(redisOptions)
	_, err := client.Ping().Result()
	if err != nil {
		panic(errors.New("redis connect fail"))
	}
	database.client = client
}