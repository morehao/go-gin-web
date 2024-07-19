package genCode

import (
	"go-gin-web/internal/genCode/config"
	"os"
)

func Run() {
	workDir, getWorkDirErr := os.Getwd()
	if getWorkDirErr != nil {
		panic(getWorkDirErr)
	}
	switch Config.CodeGen.Mode {
	case config.ModeModule:
		genModule(workDir)
	case config.ModeModel:
	case config.ModeController:
	default:
		panic("unknown mode")
	}
}
