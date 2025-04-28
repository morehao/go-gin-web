package storages

import (
	"github.com/morehao/go-tools/storages/dbmysql"
	"gorm.io/gorm"
)

func InitMultiMysql(configs []dbmysql.MysqlConfig) error {
	return dbmysql.InitMultiMysql(configs)
}

func MysqlPracticeDB() *gorm.DB {
	return dbmysql.GetDB("practice")
}
