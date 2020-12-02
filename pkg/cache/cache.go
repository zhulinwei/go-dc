package cache

import (
	"sync"

	"github.com/go-redis/redis"
	"github.com/zhulinwei/go-dc/pkg/config"
	"github.com/zhulinwei/go-dc/pkg/model"
	"github.com/zhulinwei/go-dc/pkg/util/log"
)

const (
	cacheKey = "cache"
)

type ICache interface {
	Client() *redis.Client
}

type Cache struct {
	once      sync.Once
	configs   []model.RedisConfig
	clientMap map[string]*redis.Client
}

var cache *Cache

func BuildCache() ICache {
	if cache == nil {
		cache = &Cache{configs: config.ServerConfig().Redis}
		cache.init()
	}

	return cache
}

func (cache Cache) Client() *redis.Client {
	return cache.clientMap[cacheKey]
}

func (cache Cache) init() {
	cache.once.Do(func() {
		cache.clientMap = make(map[string]*redis.Client, len(cache.configs))

		for _, redisConfig := range cache.configs {
			// 解析redis链接地址
			redisOptions, err := redis.ParseURL(redisConfig.Addr)
			if err != nil {
				log.Error("redis parse config fail", log.String("error", err.Error()))
				return
			}
			// 连接redis数据库
			client := redis.NewClient(redisOptions)
			if _, err := client.Ping().Result(); err != nil {
				log.Error("redis ping fail", log.String("error", err.Error()))
				return
			}
			// 保存mongodb客户端
			cache.clientMap[redisConfig.Name] = client
		}
	})
}
