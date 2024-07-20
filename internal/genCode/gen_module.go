package genCode

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/morehao/go-tools/gast"

	"github.com/morehao/go-tools/codeGen"
	"github.com/morehao/go-tools/gutils"
)

func genModule(workDir string) {
	cfg := Config.CodeGen.Module
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
		panic(fmt.Errorf("analysis module tpl error: %v", analysisErr))
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
			ExtraParams: ModuleExtraParams{
				PackageName:            analysisRes.PackageName,
				PackagePascalName:      analysisRes.PackagePascalName,
				ImportDirPrefix:        cfg.ImportDirPrefix,
				TableName:              analysisRes.TableName,
				Description:            cfg.Description,
				StructName:             analysisRes.StructName,
				ReceiverTypeName:       gutils.FirstLetterToLower(analysisRes.StructName),
				ReceiverTypePascalName: analysisRes.StructName,
				ApiDocTag:              cfg.ApiDocTag,
				ModuleApiPrefix:        cfg.ModuleApiPrefix,
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

	// 注册路由
	routerCallContent := fmt.Sprintf("%sRouter(routerGroup)", gutils.FirstLetterToLower(analysisRes.StructName))
	routerEnterFilepath := filepath.Join(rootDir, "/router/enter.go")
	if err := gast.AddContentToFunc(routerCallContent, "RegisterRouter", routerEnterFilepath); err != nil {
		panic(fmt.Errorf("appendContentToFunc error: %v", err))
	}
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