package config

import (
	"github.com/morehao/go-tools/glog"
	"github.com/morehao/go-tools/stores/dbmysql"
	"github.com/morehao/go-tools/stores/dbredis"
)

type Config struct {
	Server Server              `yaml:"server"`
	Log    glog.LogConfig      `yaml:"log"`
	Mysql  dbmysql.MysqlConfig `yaml:"mysql"`
	Redis  dbredis.RedisConfig `yaml:"redis"`
}

type Server struct {
	Name string `yaml:"name"` // 服务名称
	Port string `yaml:"port"` // 服务端口
	Env  string `yaml:"env"`  // 环境变量
}
