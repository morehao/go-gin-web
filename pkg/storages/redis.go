package storages

import (
	"fmt"

	"github.com/morehao/go-tools/storages/dbredis"
	"github.com/redis/go-redis/v9"
)

var (
	DemoRedis *redis.Client
)

const (
	RedisServiceNameDemo = "demo"
)

func InitMultiRedis(configs []dbredis.RedisConfig) error {
	if len(configs) == 0 {
		return nil
	}
	for _, cfg := range configs {
		client, err := dbredis.InitRedis(cfg)
		if err != nil {
			return err
		}
		switch cfg.Service {
		case RedisServiceNameDemo:
			DemoRedis = client
		default:
			return fmt.Errorf("unknown redis service: %s", cfg.Service)
		}
	}
	return nil
}
