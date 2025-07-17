package storages

import (
	"fmt"

	"github.com/morehao/go-gin-web/internal/apps/demoapp/config"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/morehao/golib/storages/dbes"
)

var (
	DemoES *elasticsearch.Client
)

const (
	ESServiceDemo = "demoapp"
)

func InitMultiEs(configs []dbes.ESConfig) error {
	if len(configs) == 0 {
		return fmt.Errorf("es config is empty")
	}
	var opts []dbes.Option
	logCfg, ok := config.Conf.Log["es"]
	if ok {
		opts = append(opts, dbes.WithLogConfig(&logCfg))
	}
	for _, cfg := range configs {
		client, _, err := dbes.InitES(&cfg, opts...)
		if err != nil {
			return err
		}
		switch cfg.Service {
		case ESServiceDemo:
			DemoES = client
		default:
			return fmt.Errorf("unknown es service name: %s", cfg.Service)
		}
	}
	return nil
}
