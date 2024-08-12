package helper

import (
	"fmt"

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
	redisClient, getRedisClientErr := dbClient.InitRedis(Config.Redis)
	if getRedisClientErr != nil {
		panic(fmt.Sprintf("get redis client error: %v", getRedisClientErr))
	}
	if redisClient == nil {
		panic(fmt.Sprintf("get redis client error: %v", getRedisClientErr))
	}
	RedisClient = redisClient
}
