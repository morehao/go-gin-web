package config

import (
	"github.com/morehao/go-tools/dbClient"
	"github.com/morehao/go-tools/glog"
)

type Config struct {
	Server Server               `yaml:"server"`
	Log    glog.LoggerConfig    `yaml:"log"`
	Mysql  dbClient.MysqlConfig `yaml:"mysql"`
	Redis  dbClient.RedisConfig `yaml:"redis"`
}

type Server struct {
	Name string `yaml:"name"` // 服务名称
	Port string `yaml:"port"` // 服务端口
	Env  string `yaml:"env"`  // 环境变量
}
