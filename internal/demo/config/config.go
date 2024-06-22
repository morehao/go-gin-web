package config

import (
	"github.com/morehao/go-tools/conf"
	"github.com/morehao/go-tools/glog"
)

var Cfg *Config

func InitConfig() {
	configFilepath := conf.GetAppRootDir() + "/config/config.yaml"
	conf.LoadConfig(configFilepath, &Cfg)
}

type Config struct {
	Server Server            `yaml:"server"`
	Log    glog.LoggerConfig `yaml:"log"`
}

type Server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}
