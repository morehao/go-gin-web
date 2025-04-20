package config

import (
	"github.com/morehao/go-tools/dbclient"
	"github.com/morehao/go-tools/glog"
)

type Config struct {
	Server Server               `yaml:"server"`
	Log    glog.LoggerConfig    `yaml:"log"`
	Mysql  dbclient.MysqlConfig `yaml:"mysql"`
	Redis  dbclient.RedisConfig `yaml:"redis"`
}

type Server struct {
	Name string `yaml:"name"` // 服务名称
	Port string `yaml:"port"` // 服务端口
	Env  string `yaml:"env"`  // 环境变量
}
