package helper

import (
	"github.com/morehao/go-tools/dbClient"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var MysqlClient *gorm.DB
var RedisClient *redis.Client

func InitDbClient() {
	mysqlClient, getMysqlClientErr := dbClient.InitMysql(Config.Mysql)
	if getMysqlClientErr != nil {
		panic("get mysql client error")
	}
	MysqlClient = mysqlClient
	RedisClient = dbClient.InitRedis(Config.Redis)
}
