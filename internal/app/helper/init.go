package helper

import (
	"go-gin-web/internal/app/config"
	"os"
	"path/filepath"

	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/glog"
)

var Config *config.Config

func SetRootDir(projectDir string) {
	if workDir, err := os.Getwd(); err != nil {
		panic("get work dir error")
	} else {
		conf.SetAppRootDir(filepath.Join(workDir, projectDir))
	}
}

func PreInit() {
	// 加载配置
	configFilepath := conf.GetAppRootDir() + "/config/config.yaml"
	var cfg config.Config
	conf.LoadConfig(configFilepath, &cfg)
	Config = &cfg

	// 初始化日志
	if err := glog.InitZapLogger(&Config.Log); err != nil {
		panic("init zap logger error")
	}
}
