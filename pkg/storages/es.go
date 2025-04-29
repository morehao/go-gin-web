package storages

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/morehao/go-tools/storages/dbes"
)

var (
	DemoES *elasticsearch.Client
)

const (
	ESServiceDemo = "demo"
)

func InitMultiEs(configs []dbes.ESConfig) error {
	for _, cfg := range configs {
		client, _, err := dbes.InitES(cfg)
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
