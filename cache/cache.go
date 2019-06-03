package cache

import (
	"sync"
	"github.com/go-redis/redis"
)

type redisCacheDataSource struct {
	redisClient   redis.UniversalClient
	servers       []string
	password      string
	db            int
}

var instance *redisCacheDataSource
var once sync.Once

func CacheDataSource() *redisCacheDataSource {
	once.Do(func() {
		instance = &redisCacheDataSource{}
	})
	return instance
}

func GetRedisClient() redis.UniversalClient {
	return instance.redisClient
}

func (p *redisCacheDataSource) DataSourceInit(servers []string, password string, db int) {
	p.redisClient = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    servers,
		Password: password,
		DB:       db,
	})
	p.password = password
	p.servers = servers
	p.db = db
}

