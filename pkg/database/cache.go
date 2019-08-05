package database

import (
	"errors"
	"github.com/go-redis/redis"
)

type Cache struct {}

func (cache *Cache) InitRedis(redisConfig *RedisConfig) *redis.Client {
	var redisOptions = &redis.Options{
		Addr: redisConfig.Url,
	}
	client := redis.NewClient(redisOptions)
	_, err := client.Ping().Result()
	if err != nil {
		panic(errors.New("redis connect fail"))
	}
	return client
}
