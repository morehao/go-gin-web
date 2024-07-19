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
	Mysql   dbClient.MysqlConfig `yaml:"mysql"`
	Log     glog.LoggerConfig    `yaml:"log"`
	CodeGen CodeGen              `yaml:"code_gen"`
}

type CodeGen struct {
	Mode   string       `yaml:"mode"`   // 生成模式，支持：module、model、controller
	Module ModuleConfig `yaml:"module"` // 模块名称
}

type ModuleConfig struct {
	TplDir            string `yaml:"tpl_dir"`            // 模板目录
	RootDir           string `yaml:"root_dir"`           // 项目内当前项目的根目录(如internal/genCode)
	ImportDirPrefix   string `yaml:"import_dir_prefix"`  // import目录前缀
	ModuleDescription string `yaml:"module_description"` // 模块描述
	ApiDocTag         string `yaml:"api_doc_tag"`        // api文档tag
	ModuleApiPrefix   string `yaml:"module_api_prefix"`  // api前缀
	PackageName       string `yaml:"package_name"`       // 包名
	TableName         string `yaml:"table_name"`         // 表名
}

const (
	ModeModule     = "module"
	ModeModel      = "model"
	ModeController = "controller"
)
