package storages

import (
	"fmt"

	"github.com/morehao/go-gin-web/internal/apps/demoapp/config"
  "github.com/morehao/golib/dbstore/dbredis"
	"github.com/redis/go-redis/v9"
)

var (
	DemoRedis *redis.Client
)

const (
	RedisServiceNameDemo = "demoapp"
)

func InitMultiRedis(configs []dbredis.RedisConfig) error {
	if len(configs) == 0 {
		return fmt.Errorf("redis config is empty")
	}
	var opts []dbredis.Option
	logCfg, ok := config.Conf.Log["redis"]
	if ok {
		opts = append(opts, dbredis.WithLogConfig(&logCfg))
	}
	for _, cfg := range configs {
		client, err := dbredis.InitRedis(&cfg, opts...)
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
