package main

import (
	"fmt"

	"go-gin-web/apps/demo/config"
	"go-gin-web/pkg/storages"

	"github.com/morehao/go-tools/glog"
)

func serverInit() error {
	if err := preInit(); err != nil {
		return err
	}
	if err := resourceInit(); err != nil {
		return err
	}
	return nil
}

func preInit() error {
	config.InitConf()
	if err := glog.InitLogger(&config.Conf.Log); err != nil {
		return fmt.Errorf("init logger failed: " + err.Error())
	}
	return nil
}

func resourceInit() error {
	if err := storages.InitMultiMysql(config.Conf.MysqlConfigs); err != nil {
		return fmt.Errorf("init mysql failed: " + err.Error())
	}
	if err := storages.InitMultiRedis(config.Conf.RedisConfigs); err != nil {
		return fmt.Errorf("init redis failed: " + err.Error())
	}
	return nil
}
