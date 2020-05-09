package cache

import (
	"log"
	"sync"

	"github.com/zhulinwei/go-dc/pkg/model"

	"github.com/go-redis/redis"
	"github.com/zhulinwei/go-dc/pkg/config"
)

const (
	cacheKey = "cache"
)

type ICache interface {
	Client() *redis.Client
}

type Cache struct {
	once      sync.Once
	configs   []model.ReidsConfig
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
		var wg sync.WaitGroup
		length := len(cache.configs)
		wg.Add(length)
		cache.clientMap = make(map[string]*redis.Client, length)

		for _, redisConfig := range cache.configs {

			go func(config model.ReidsConfig, wg *sync.WaitGroup) {
				defer wg.Done()

				// 解析redis链接地址
				redisOptions, err := redis.ParseURL(redisConfig.Addr)
				if err != nil {
					log.Printf("redis parse config failed: %v", err.Error())
					return
				}
				// 连接redis数据库
				client := redis.NewClient(redisOptions)
				if _, err := client.Ping().Result(); err != nil {
					log.Printf("redis ping failed: %v", err.Error())
					return
				}
				// 保存mongodb客户端
				cache.clientMap[redisConfig.Name] = client
			}(redisConfig, &wg)
		}

		wg.Wait()
	})
}
