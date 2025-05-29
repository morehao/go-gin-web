package storages

import (
	"fmt"

	"go-gin-web/internal/apps/demoapp/config"

	"github.com/morehao/golib/storages/dbmysql"
	"gorm.io/gorm"
)

var (
	DBDemo *gorm.DB
)

const (
	DBNameDemo = "demo"
)

func InitMultiMysql(configs []dbmysql.MysqlConfig) error {
	if len(configs) == 0 {
		return fmt.Errorf("mysql config is empty")
	}

	var opts []dbmysql.Option
	logCfg, ok := config.Conf.Log["gorm"]
	if ok {
		opts = append(opts, dbmysql.WithLogConfig(&logCfg))
	}
	for _, cfg := range configs {
		client, err := dbmysql.InitMysql(&cfg, opts...)
		if err != nil {
			return fmt.Errorf("init mysql failed: " + err.Error())
		}
		switch cfg.Database {
		case DBNameDemo:
			DBDemo = client
		default:
			return fmt.Errorf("unknown database: " + cfg.Database)
		}
	}
	return nil
}
