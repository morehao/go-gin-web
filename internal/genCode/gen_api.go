package genCode

import (
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/morehao/go-tools/codeGen"
	"github.com/morehao/go-tools/gutils"
)

func genApi(workDir string) {
	cfg := Config.CodeGen.Api
	tplDir := filepath.Join(workDir, cfg.TplDir)
	rootDir := filepath.Join(workDir, cfg.RootDir)
	layerDirMap := map[codeGen.LayerName]string{
		codeGen.LayerNameErrorCode: filepath.Join(rootDir, "/pkg"),
	}
	analysisCfg := &codeGen.ApiCfg{
		CommonConfig: codeGen.CommonConfig{
			TplDir:      tplDir,
			PackageName: cfg.PackageName,
			RootDir:     rootDir,
			LayerDirMap: layerDirMap,
		},
		TargetFilename: cfg.TargetFilename,
	}
	gen := codeGen.NewGenerator()
	analysisRes, analysisErr := gen.AnalysisApiTpl(analysisCfg)
	if analysisErr != nil {
		panic(fmt.Errorf("analysis api tpl error: %v", analysisErr))
	}
	receiverTypePascalName := gutils.SnakeToPascal(gutils.TrimFileExtension(cfg.TargetFilename))
	var genParamsList []codeGen.GenParamsItem
	for _, v := range analysisRes.TplAnalysisList {

		genParamsList = append(genParamsList, codeGen.GenParamsItem{
			TargetDir:      v.TargetDir,
			TargetFileName: v.TargetFilename,
			Template:       v.Template,
			ExtraParams: ApiExtraParams{
				PackageName:            analysisRes.PackageName,
				PackagePascalName:      analysisRes.PackagePascalName,
				ImportDirPrefix:        cfg.ImportDirPrefix,
				TargetFileExist:        v.TargetFileExist,
				Description:            cfg.Description,
				ReceiverTypeName:       gutils.FirstLetterToLower(receiverTypePascalName),
				ReceiverTypePascalName: receiverTypePascalName,
				HttpMethod:             cfg.HttpMethod,
				FunctionName:           gutils.FirstLetterToUpper(cfg.FunctionName),
				ApiDocTag:              cfg.ApiDocTag,
				ApiPath:                cfg.ApiPath,
				Template:               v.Template,
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

type ApiExtraParams struct {
	ImportDirPrefix        string
	PackageName            string
	PackagePascalName      string
	Description            string
	TargetFileExist        bool
	HttpMethod             string
	FunctionName           string
	ReceiverTypeName       string
	ReceiverTypePascalName string
	ApiPath                string
	ApiDocTag              string
	Template               *template.Template
}
