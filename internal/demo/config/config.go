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
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
