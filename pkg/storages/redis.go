package storages

import (
	"github.com/morehao/go-tools/storages/dbredis"
	"github.com/redis/go-redis/v9"
)

func InitMultiRedis(configs []dbredis.RedisConfig) error {
	return dbredis.InitMultiRedis(configs)
}

func RedisClient() *redis.Client {
	return dbredis.GetClient("go-gin-web")
}
