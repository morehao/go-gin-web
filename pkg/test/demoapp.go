package test

import (
	"fmt"
	"path"
	"runtime"

	"go-gin-web/internal/apps/demoapp/config"
	"go-gin-web/pkg/storages"

	"github.com/morehao/golib/glog"
)

type demo struct {
}

func newDemo() Initializer {
	return &demo{}
}

func (d *demo) Initialize() error {
	d.initConf()
	if err := d.preInit(); err != nil {
		return err
	}
	if err := d.resourceInit(); err != nil {
		return err
	}
	return nil
}

func (d *demo) initConf() {
	_, file, _, _ := runtime.Caller(0)
	rootDir := path.Join(path.Dir(path.Dir(path.Dir(file))), "apps", AppNameDemo)
	config.SetRootDir(rootDir)
	config.InitConf()
	return
}

func (d *demo) preInit() error {
	config.InitConf()
	defaultLogCfg := config.Conf.Log["default"]
	if err := glog.InitLogger(&defaultLogCfg); err != nil {
		return fmt.Errorf("init logger failed: " + err.Error())
	}
	return nil
}

func (d *demo) resourceInit() error {
	if err := storages.InitMultiMysql(config.Conf.MysqlConfigs); err != nil {
		return fmt.Errorf("init mysql failed: " + err.Error())
	}
	if err := storages.InitMultiRedis(config.Conf.RedisConfigs); err != nil {
		return fmt.Errorf("init redis failed: " + err.Error())
	}
	if err := storages.InitMultiEs(config.Conf.ESConfigs); err != nil {
		return fmt.Errorf("init es failed: " + err.Error())
	}
	return nil
}

func (d *demo) Close() error {
	glog.Close()
	return nil
}
