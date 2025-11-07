package demoapp

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/morehao/go-gin-web/apps/demoapp/config"
	_ "github.com/morehao/go-gin-web/apps/demoapp/docs"
	"github.com/morehao/go-gin-web/apps/demoapp/middleware"
	"github.com/morehao/go-gin-web/apps/demoapp/router"
	"github.com/morehao/go-gin-web/pkg/storages"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/glog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	_, workDir, _, _ := runtime.Caller(0)
	rootDir := filepath.Dir(workDir)
	config.SetRootDir(rootDir)
	if err := serverInit(); err != nil {
		panic(fmt.Sprintf("server init failed, error: %v", err))
	}
	if config.Conf.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	defer glog.Close()

	engine := gin.New()
	engine.Use(gin.Recovery())
	routerGroup := engine.Group(fmt.Sprintf("/%s", config.Conf.Server.Name))
	if config.Conf.Server.Env == "dev" {
		routerGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	routerGroup.Use(middleware.AccessLog())
	router.RegisterRouter(routerGroup)
	if err := engine.Run(fmt.Sprintf(":%s", config.Conf.Server.Port)); err != nil {
		fmt.Println(fmt.Sprintf("%s run fail, port:%s", config.Conf.Server.Name, config.Conf.Server.Port))
	} else {
		fmt.Println(fmt.Sprintf("%s run success, port:%s", config.Conf.Server.Name, config.Conf.Server.Port))
	}
}

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
	defaultLogCfg := config.Conf.Log["default"]
	if err := glog.InitLogger(&defaultLogCfg); err != nil {
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
	if err := storages.InitMultiEs(config.Conf.ESConfigs); err != nil {
		return fmt.Errorf("init es failed: " + err.Error())
	}
	return nil
}
