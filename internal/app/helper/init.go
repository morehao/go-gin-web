package helper

import (
	"fmt"
	"go-gin-web/internal/app/config"

	"github.com/morehao/go-tools/dbClient"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	"github.com/morehao/go-tools/glog"
	"go.uber.org/zap"

	"github.com/morehao/go-tools/conf"
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
	if err := glog.NewLogger(&Config.Log, glog.WithZapOptions(zap.AddCallerSkip(3))); err != nil {
		panic("init zap logger error")
	}
}

func ResourceInit() {
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
