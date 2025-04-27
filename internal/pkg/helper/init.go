package helper

import (
	"fmt"

	"go-gin-web/config"

	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/glog"
	"github.com/morehao/go-tools/stores/dbmysql"
	"github.com/morehao/go-tools/stores/dbredis"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var Config *config.Config

var MysqlClient *gorm.DB
var RedisClient *redis.Client

func SetRootDir(rootDir string) {
	conf.SetAppRootDir(rootDir)
}

func PreInit() {
	ConfInit()
	LogInit()
}

func ConfInit() {
	// 加载配置
	configFilepath := conf.GetAppRootDir() + "/config/config.yaml"
	var cfg config.Config
	conf.LoadConfig(configFilepath, &cfg)
	Config = &cfg
}

func LogInit() {
	// 初始化日志
	if err := glog.InitLogger(&Config.Log, glog.WithCallerSkip(3)); err != nil {
		panic("init zap logger error")
	}
}

func ResourceInit() {
	mysqlClient, getMysqlClientErr := dbmysql.InitMysql(Config.Mysql)
	if getMysqlClientErr != nil {
		panic("get mysql client error")
	}
	MysqlClient = mysqlClient
	redisClient, getRedisClientErr := dbredis.InitRedis(Config.Redis)
	if getRedisClientErr != nil {
		panic(fmt.Sprintf("get redis client error: %v", getRedisClientErr))
	}
	if redisClient == nil {
		panic(fmt.Sprintf("get redis client error: %v", getRedisClientErr))
	}
	RedisClient = redisClient
}
