package env

import (
	"github.com/gin-gonic/gin"
	"go-practict/utils"
	"os"
	"path/filepath"
)

const DefaultRootPath = "."

const (
	DockerRunEnv   = "RUN_ENV"
	ServiceNameEnv = "SERVICE_NAME"
)

const (
	RunEnvTest   = 0
	RunEnvPre    = 1
	RunEnvOnline = 2
)

var (
	LocalIP string
	AppName string
	RunMode string

	ServiceName string

	runEnv int

	rootPath        string
	dockerPlateForm bool
)

func init() {
	LocalIP = utils.GetLocalIp()
	// 运行环境
	RunMode = gin.ReleaseMode
	r := os.Getenv(DockerRunEnv)
	switch r {
	case "prod":
		runEnv = RunEnvOnline
	case "pre":
		runEnv = RunEnvPre
	default:
		runEnv = RunEnvTest
		RunMode = gin.DebugMode
	}

	ServiceName = os.Getenv(ServiceNameEnv)

	gin.SetMode(RunMode)
}

func SetAppName(appName string) {
	if !dockerPlateForm {
		AppName = appName
	}
}

func GetAppName() string {
	return AppName
}

// SetRootPath 设置应用的根目录
func SetRootPath(r string) {
	if !dockerPlateForm {
		rootPath = r
	}
}

// GetRootPath 返回应用的根目录
func GetRootPath() string {
	if rootPath != "" {
		return rootPath
	} else {
		return DefaultRootPath
	}
}

// GetConfDirPath 返回配置文件目录绝对地址
func GetConfDirPath() string {
	return filepath.Join(GetRootPath(), "conf")
}

// GetLogDirPath 返回log目录的绝对地址
func GetLogDirPath() string {
	return filepath.Join(GetRootPath(), "log")
}

func GetRunEnv() int {
	return runEnv
}
