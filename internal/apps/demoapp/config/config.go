package config

import (
	"fmt"
	"os"

	"github.com/morehao/golib/conf"
	"github.com/morehao/golib/dbstore/dbes"
	"github.com/morehao/golib/dbstore/dbmysql"
	"github.com/morehao/golib/dbstore/dbredis"
	"github.com/morehao/golib/glog"
	"github.com/morehao/golib/protocol/gresty"
)

var Conf *Config

type Config struct {
	Server       Server                    `yaml:"server"`
	Log          map[string]glog.LogConfig `yaml:"log"`
	MysqlConfigs []dbmysql.MysqlConfig     `yaml:"mysql_configs"`
	RedisConfigs []dbredis.RedisConfig     `yaml:"redis_configs"`
	ESConfigs    []dbes.ESConfig           `yaml:"es_configs"`
	Client       Client                    `yaml:"client"`
}

type Server struct {
	Name string `yaml:"name"` // 服务名称
	Port string `yaml:"port"` // 服务端口
	Env  string `yaml:"env"`  // 环境变量
}

type Client struct {
	HTTPBingo *gresty.Client `yaml:"httpbingo"`
}

func SetRootDir(rootDir string) {
	conf.SetAppRootDir(rootDir)
}

func InitConf() {
	// 读取环境变量，如果没设置，则用默认路径
	configPath := os.Getenv("APP_CONFIG_PATH")
	if configPath == "" {
		configPath = conf.GetAppRootDir() + "/config/config.yaml"
	}
	fmt.Println("Load config file:", configPath)

	var cfg Config
	conf.LoadConfig(configPath, &cfg)
	Conf = &cfg
}
