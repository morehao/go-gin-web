package main

import (
	"go-gin-web/internal/demo/helper"
	"os"
	"path/filepath"
	"text/template"

	"github.com/morehao/go-tools/gutils"

	"github.com/morehao/go-tools/codeGen"
)

func main() {
	// 初始化配置
	helper.PreInit()
	helper.InitResource()
	workDir, getWorkDirErr := os.Getwd()
	if getWorkDirErr != nil {
		panic(getWorkDirErr)
	}
	genModule(workDir)
	// 读取配置文件
	// 生成代码
}

func genModule(workDir string) {
	tplDir := filepath.Join(workDir, "internal/resource/codeTpl/module")
	rootDir := filepath.Join(workDir, "internal/demo")
	layerDirMap := map[codeGen.LayerName]string{
		codeGen.LayerNameErrorCode: filepath.Join(rootDir, "/pkg"),
	}
	importDirPrefix := "go-gin-web/internal/demo"
	moduleDescription := "用户"
	apiDocTag := "用户管理"
	moduleApiPrefix := "/go-gin-web/user"
	cfg := &codeGen.ModuleCfg{
		TplDir:      tplDir,
		PackageName: "user",
		TableName:   "user",
		RootDir:     rootDir,
		LayerDirMap: layerDirMap,
	}
	gen := codeGen.NewGenerator()
	tplParamsRes, getModuleParamErr := gen.GetModuleTemplateParams(helper.MysqlClient, cfg)
	if getModuleParamErr != nil {
		panic(getModuleParamErr)
	}

	type ModelField struct {
		FieldName    string // 字段名称
		FieldType    string // 字段数据类型，如int、string
		ColumnName   string // 列名
		ColumnType   string // 列数据类型，如varchar(255)
		Comment      string // 字段注释
		IsPrimaryKey bool   // 是否是主键
	}

	type ModuleExtraParams struct {
		ImportDirPrefix        string
		PackageName            string
		PackagePascalName      string
		TableName              string
		Description            string
		StructName             string
		ReceiverTypeName       string
		ReceiverTypePascalName string
		ModuleApiPrefix        string
		ApiDocTag              string
		Template               *template.Template
		ModelFields            []ModelField
	}
	var genParamsList []codeGen.GenParamsItem
	for _, v := range tplParamsRes.TemplateList {
		var modelFields []ModelField
		for _, field := range v.ModelFields {
			modelFields = append(modelFields, ModelField{
				FieldName:    field.FieldName,
				FieldType:    field.FieldType,
				ColumnName:   field.ColumnName,
				ColumnType:   field.ColumnType,
				Comment:      field.Comment,
				IsPrimaryKey: field.ColumnKey == codeGen.ColumnKeyPRI,
			})
		}

		genParamsList = append(genParamsList, codeGen.GenParamsItem{
			TargetDir:      v.TargetDir,
			TargetFileName: v.TargetFilename,
			Template:       v.Template,
			ExtraParams: ModuleExtraParams{
				PackageName:            tplParamsRes.PackageName,
				PackagePascalName:      tplParamsRes.PackagePascalName,
				ImportDirPrefix:        importDirPrefix,
				TableName:              tplParamsRes.TableName,
				Description:            moduleDescription,
				StructName:             tplParamsRes.StructName,
				ReceiverTypeName:       gutils.FirstLetterToLower(tplParamsRes.StructName),
				ReceiverTypePascalName: gutils.FirstLetterToUpper(tplParamsRes.StructName),
				ApiDocTag:              apiDocTag,
				ModuleApiPrefix:        moduleApiPrefix,
				Template:               v.Template,
				ModelFields:            modelFields,
			},
		})

	}
	genParams := &codeGen.GenParams{
		ParamsList: genParamsList,
	}
	if err := gen.Gen(genParams); err != nil {
		panic(err)
	}
}
