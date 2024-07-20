package genCode

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/morehao/go-tools/codeGen"
)

func genModel(workDir string) {
	cfg := Config.CodeGen.Model
	tplDir := filepath.Join(workDir, cfg.TplDir)
	rootDir := filepath.Join(workDir, cfg.RootDir)
	layerDirMap := map[codeGen.LayerName]string{
		codeGen.LayerNameErrorCode: filepath.Join(rootDir, "/pkg"),
	}
	analysisCfg := &codeGen.ModuleCfg{
		CommonConfig: codeGen.CommonConfig{
			TplDir:      tplDir,
			PackageName: cfg.PackageName,
			RootDir:     rootDir,
			LayerDirMap: layerDirMap,
		},
		TableName: cfg.TableName,
	}
	gen := codeGen.NewGenerator()
	analysisRes, analysisErr := gen.AnalysisModuleTpl(MysqlClient, analysisCfg)
	if analysisErr != nil {
		panic(fmt.Errorf("analysis model tpl error: %v", analysisErr))
	}

	var genParamsList []codeGen.GenParamsItem
	for _, v := range analysisRes.TplAnalysisList {
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
			ExtraParams: ModelExtraParams{
				PackageName:       analysisRes.PackageName,
				PackagePascalName: analysisRes.PackagePascalName,
				ImportDirPrefix:   cfg.ImportDirPrefix,
				TableName:         analysisRes.TableName,
				Description:       cfg.Description,
				StructName:        analysisRes.StructName,
				Template:          v.Template,
				ModelFields:       modelFields,
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

type ModelField struct {
	FieldName    string // 字段名称
	FieldType    string // 字段数据类型，如int、string
	ColumnName   string // 列名
	ColumnType   string // 列数据类型，如varchar(255)
	Comment      string // 字段注释
	IsPrimaryKey bool   // 是否是主键
}

type ModelExtraParams struct {
	ImportDirPrefix   string
	PackageName       string
	PackagePascalName string
	TableName         string
	Description       string
	StructName        string
	Template          *template.Template
	ModelFields       []ModelField
}