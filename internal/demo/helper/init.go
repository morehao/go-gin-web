package helper

import (
	"go-gin-web/internal/demo/config"
	"os"
	"path/filepath"

	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/gcore"
	"github.com/morehao/go-tools/glog"
	"gorm.io/gorm"
)

var MysqlClient *gorm.DB

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
	mysqlClient, getMysqlClientErr := gcore.InitMysql(config.Cfg.Mysql)
	if getMysqlClientErr != nil {
		panic("get mysql client error")
	}
	MysqlClient = mysqlClient
}

func Clear() {
	glog.Close()
}
