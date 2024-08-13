package genCode

import (
	"fmt"
	"io/fs"
	"text/template"

	"github.com/morehao/go-tools/codeGen"
)

func Run() {
	// 列出嵌入文件系统中的所有文件
	templatesFS, err := codeGen.NewGenerator().GetTplFs(codeGen.FrameworkGin)
	if err != nil {
		panic(err)
	}
	err = fs.WalkDir(templatesFS, "ginTpl/api", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			fmt.Println("File:", path)
		} else {
			tplInst := template.New(path)
			_, err = tplInst.ParseFS(templatesFS, path+"/*")
			fmt.Println("Directory:", path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	// workDir, getWorkDirErr := os.Getwd()
	// if getWorkDirErr != nil {
	// 	panic(getWorkDirErr)
	// }
	// switch Config.CodeGen.Mode {
	// case config.ModeModule:
	// 	genModule(workDir)
	// case config.ModeModel:
	// 	genModel(workDir)
	// case config.ModeApi:
	// 	genApi(workDir)
	// default:
	// 	panic("unknown mode")
	// }
	fmt.Println(Config.CodeGen.Mode, "生成模式下，生成代码完成")
}
