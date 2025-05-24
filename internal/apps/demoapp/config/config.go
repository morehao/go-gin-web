package config

import (
	"github.com/morehao/golib/conf"
	"github.com/morehao/golib/glog"
	"github.com/morehao/golib/storages/dbes"
	"github.com/morehao/golib/storages/dbmysql"
	"github.com/morehao/golib/storages/dbredis"
)

var Conf *Config

type Config struct {
	Server       Server                    `yaml:"server"`
	Log          map[string]glog.LogConfig `yaml:"log"`
	MysqlConfigs []dbmysql.MysqlConfig     `yaml:"mysql_configs"`
	RedisConfigs []dbredis.RedisConfig     `yaml:"redis_configs"`
	ESConfigs    []dbes.ESConfig           `yaml:"es_configs"`
}

type Server struct {
	Name string `yaml:"name"` // 服务名称
	Port string `yaml:"port"` // 服务端口
	Env  string `yaml:"env"`  // 环境变量
}

func SetRootDir(rootDir string) {
	conf.SetAppRootDir(rootDir)
}

func InitConf() {
	// 加载配置
	configFilepath := conf.GetAppRootDir() + "/config/config.yaml"
	var cfg Config
	conf.LoadConfig(configFilepath, &cfg)
	Conf = &cfg
}
