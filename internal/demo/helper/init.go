package helper

import (
	"go-gin-web/internal/demo/config"
	"os"
	"path/filepath"

	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/glog"
)

var Config *config.Config

func init() {
	if workDir, err := os.Getwd(); err != nil {
		panic("get work dir error")
	} else {
		conf.SetAppRootDir(filepath.Join(workDir, "/internal/demo"))
	}
	Config = config.InitConfig()
	if err := glog.InitZapLogger(&Config.Log); err != nil {
		panic("init zap logger error")
	}
}
