package storages

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/morehao/go-tools/storages/dbes"
)

func InitMultiEs(configs []dbes.ESConfig) error {
	return dbes.InitMultiES(configs)
}

func EsClient() *elasticsearch.Client {
	return dbes.GetSimpleClient("go-gin-web")
}
