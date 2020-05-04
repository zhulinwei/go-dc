package cache

import (
	"log"
	"sync"

	"github.com/go-redis/redis"
	"github.com/zhulinwei/go-dc/pkg/config"
)

var redisOnce sync.Once
var redisMutex sync.Mutex
var redisClientMap map[string]*redis.Client

type ICache interface {
	Client() *redis.Client
}

type Cache struct {
	ClientMap map[string]*redis.Client
}

func BuildCache() Cache {
	initCache()
	return Cache{
		ClientMap: redisClientMap,
	}
}

func (cache Cache) Client() *redis.Client {
	return cache.ClientMap["cache"]
}

func initCache() {
	redisConfigs := config.ServerConfig().Redis
	redisOnce.Do(func() {
		redisMutex.Lock()
		defer redisMutex.Unlock()

		redisClientMap = make(map[string]*redis.Client, len(redisConfigs))
		for _, redisConfig := range redisConfigs {
			var err error
			var redisOptions *redis.Options

			// 解析redis链接地址
			if redisOptions, err = redis.ParseURL(redisConfig.Addr); err != nil {
				log.Printf("redis parse config failed: %v", err)
				return
			}

			// 连接redis数据库
			client := redis.NewClient(redisOptions)
			if _, err := client.Ping().Result(); err != nil {
				log.Printf("redis ping failed: %v", err)
				return
			}

			// 保存mongodb客户端
			redisClientMap[redisConfig.Name] = client
		}
	})
}
