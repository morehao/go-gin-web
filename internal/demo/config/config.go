package config

import (
	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/dbClient"
	"github.com/morehao/go-tools/glog"
)

func InitConfig() *Config {
	var cfg Config
	configFilepath := conf.GetAppRootDir() + "/config/config.yaml"
	conf.LoadConfig(configFilepath, &cfg)
	return &cfg
}

type Config struct {
	Server Server               `yaml:"server"`
	Log    glog.LoggerConfig    `yaml:"log"`
	Mysql  dbClient.MysqlConfig `yaml:"mysql"`
	Redis  dbClient.RedisConfig `yaml:"redis"`
}

type Server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
