package genCode

import (
	"path/filepath"
	"text/template"

	"github.com/morehao/go-tools/codeGen"
	"github.com/morehao/go-tools/gutils"
)

func genModule(workDir string) {
	tplDir := filepath.Join(workDir, "internal/resource/codeTpl/module")
	rootDir := filepath.Join(workDir, "internal/demo")
	layerDirMap := map[codeGen.LayerName]string{
		codeGen.LayerNameErrorCode: filepath.Join(rootDir, "/pkg"),
	}
	cfg := &codeGen.ModuleCfg{
		TplDir:      tplDir,
		PackageName: Config.CodeGen.Module.PackageName,
		TableName:   Config.CodeGen.Module.TableName,
		RootDir:     rootDir,
		LayerDirMap: layerDirMap,
	}
	gen := codeGen.NewGenerator()
	tplParamsRes, getModuleParamErr := gen.GetModuleTemplateParams(MysqlClient, cfg)
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
				ImportDirPrefix:        Config.CodeGen.Module.ImportDirPrefix,
				TableName:              tplParamsRes.TableName,
				Description:            Config.CodeGen.Module.ModuleDescription,
				StructName:             tplParamsRes.StructName,
				ReceiverTypeName:       gutils.FirstLetterToLower(tplParamsRes.StructName),
				ReceiverTypePascalName: tplParamsRes.StructName,
				ApiDocTag:              Config.CodeGen.Module.ApiDocTag,
				ModuleApiPrefix:        Config.CodeGen.Module.ModuleApiPrefix,
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
