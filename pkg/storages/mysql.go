package storages

import (
	"fmt"

	"github.com/morehao/go-tools/storages/dbmysql"
	"gorm.io/gorm"
)

var (
	DBDemo *gorm.DB
)

const (
	DBNameDemo = "practice"
)

func InitMultiMysql(configs []dbmysql.MysqlConfig) error {
	if len(configs) == 0 {
		return fmt.Errorf("mysql config is empty")
	}
	for _, cfg := range configs {
		client, err := dbmysql.InitMysql(cfg)
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
