package helper

import (
	"go-gin-web/internal/demo/config"
	"os"
	"path/filepath"

	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/dbClient"
	"github.com/morehao/go-tools/glog"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var MysqlClient *gorm.DB
var RedisClient *redis.Client

func PreInit() {
	if workDir, err := os.Getwd(); err != nil {
		panic("get work dir error")
	} else {
		conf.SetAppRootDir(filepath.Join(workDir, "/internal/demo"))
	}
	config.InitConfig()
	if err := glog.InitZapLogger(&config.Cfg.Log); err != nil {
		panic("init zap logger error")
	}
}

func InitResource() {
	mysqlClient, getMysqlClientErr := dbClient.InitMysql(config.Cfg.Mysql)
	if getMysqlClientErr != nil {
		panic("get mysql client error")
	}
	MysqlClient = mysqlClient
	RedisClient = dbClient.InitRedis(config.Cfg.Redis)
}

func Clear() {
	glog.Close()
}
