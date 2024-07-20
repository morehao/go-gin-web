package genCode

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/morehao/go-tools/gast"

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
	receiverTypeName := gutils.FirstLetterToLower(receiverTypePascalName)
	var genParamsList []codeGen.GenParamsItem
	var isNewRouter bool
	var controllerFilepath, serviceFilepath string
	for _, v := range analysisRes.TplAnalysisList {
		switch v.LayerName {
		case codeGen.LayerNameRouter:
			isNewRouter = !v.TargetFileExist
		case codeGen.LayerNameController:
			controllerFilepath = filepath.Join(v.TargetDir, v.TargetFilename)
		case codeGen.LayerNameService:
			serviceFilepath = filepath.Join(v.TargetDir, v.TargetFilename)
		}
		if v.LayerName == codeGen.LayerNameRouter {
			isNewRouter = !v.TargetFileExist
		}

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
				ReceiverTypeName:       receiverTypeName,
				ReceiverTypePascalName: receiverTypePascalName,
				HttpMethod:             cfg.HttpMethod,
				FunctionName:           gutils.FirstLetterToUpper(cfg.FunctionName),
				ApiDocTag:              cfg.ApiDocTag,
				ApiPrefix:              strings.TrimSuffix(cfg.ApiPrefix, "/"),
				ApiSuffix:              strings.TrimLeft(cfg.ApiSuffix, "/"),
				ApiGroup:               cfg.ApiGroup,
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

	// 将方法添加到interface接口中
	controllerInterfaceName := fmt.Sprintf("%sCtr", receiverTypePascalName)
	if err := gast.AddMethodToInterfaceInFile(controllerFilepath, controllerInterfaceName, receiverTypeName+"Ctr", cfg.FunctionName); err != nil {
		panic(fmt.Errorf("add controller method to interface error: %w", err))
	}
	serviceInterfaceName := fmt.Sprintf("%Svc", receiverTypePascalName)
	if err := gast.AddMethodToInterfaceInFile(serviceFilepath, serviceInterfaceName, receiverTypeName+"Svc", cfg.FunctionName); err != nil {
		panic(fmt.Errorf("add service method to interface error: %w", err))
	}

	// 	注册路由
	if isNewRouter {
		routerCallContent := fmt.Sprintf("%sRouter(routerGroup)", receiverTypeName)
		routerEnterFilepath := filepath.Join(rootDir, "/router/enter.go")
		if err := gast.AddContentToFunc(routerCallContent, "RegisterRouter", routerEnterFilepath); err != nil {
			panic(fmt.Errorf("appendContentToFunc error: %v", err))
		}
	} else {
		routerCallContent := fmt.Sprintf(`routerGroup.%s("/%s", %sCtr.%s)`, cfg.HttpMethod, cfg.ApiSuffix, receiverTypeName, cfg.FunctionName)
		routerEnterFilepath := filepath.Join(rootDir, fmt.Sprintf("/router/%s.go", gutils.TrimFileExtension(cfg.TargetFilename)))
		if err := gast.AddContentToFunc(routerCallContent, fmt.Sprintf("%sRouter", receiverTypeName), routerEnterFilepath); err != nil {
			panic(fmt.Errorf("appendContentToFunc error: %v", err))
		}
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
	ApiGroup               string
	ApiPrefix              string
	ApiSuffix              string
	ApiDocTag              string
	Template               *template.Template
}
