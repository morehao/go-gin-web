package genCode

import (
	"go-gin-web/internal/genCode/config"
	"os"
	"path/filepath"

	"github.com/morehao/go-tools/glog"

	"gorm.io/gorm"

	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/dbClient"
)

var Config *config.Config
var MysqlClient *gorm.DB

func init() {
	// 初始化配置
	if workDir, err := os.Getwd(); err != nil {
		panic("get work dir error")
	} else {
		conf.SetAppRootDir(filepath.Join(workDir, "/internal/genCode"))
	}
	Config = config.InitConfig()
	if err := glog.InitZapLogger(&Config.Log); err != nil {
		panic("init zap logger error")
	}
	mysqlClient, getMysqlClientErr := dbClient.InitMysql(Config.Mysql)
	if getMysqlClientErr != nil {
		panic("get mysql client error")
	}
	MysqlClient = mysqlClient
}
