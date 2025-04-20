package helper

import (
	"fmt"

	"go-gin-web/config"

	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/dbclient"
	"github.com/morehao/go-tools/glog"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
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
	if err := glog.NewLogger(&Config.Log, glog.WithZapOptions(zap.AddCallerSkip(3))); err != nil {
		panic("init zap logger error")
	}
}

func ResourceInit() {
	mysqlClient, getMysqlClientErr := dbclient.InitMysql(Config.Mysql)
	if getMysqlClientErr != nil {
		panic("get mysql client error")
	}
	MysqlClient = mysqlClient
	redisClient, getRedisClientErr := dbclient.InitRedis(Config.Redis)
	if getRedisClientErr != nil {
		panic(fmt.Sprintf("get redis client error: %v", getRedisClientErr))
	}
	if redisClient == nil {
		panic(fmt.Sprintf("get redis client error: %v", getRedisClientErr))
	}
	RedisClient = redisClient
}
