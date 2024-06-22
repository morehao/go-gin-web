package helper

import (
	"go-gin-web/internal/demo/config"
	"os"
	"path/filepath"

	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/glog"
)

func PreInit() {
	if workDir, err := os.Getwd(); err != nil {
		panic("get work dir error")
	} else {
		conf.SetAppRootDir(filepath.Join(workDir, "/internal/demo"))
	}
	config.InitConfig()
	glog.InitZapLogger(&config.Cfg.Log)
}

func Clear() {
	glog.Close()
}
