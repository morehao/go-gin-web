package router

import (
	"{{.ImportDirPrefix}}/controller/ctr{{.PackagePascalName}}"

	"github.com/gin-gonic/gin"
)
{{if not .TargetFileExist}}
// {{.ReceiverTypeName}}Router 初始化{{.Description}}路由信息
func {{.ReceiverTypeName}}Router(privateRouter *gin.RouterGroup) {
	{{.ReceiverTypeName}}Ctr := ctr{{.PackagePascalName}}.New{{.ReceiverTypePascalName}}Ctr()
	routerGroup := privateRouter.Group("{{.ApiGroup}}")
	{
		routerGroup.{{.HttpMethod}}("{{.ApiSuffix}}", {{.ReceiverTypeName}}Ctr.{{.FunctionName}})   // {{.Description}}
	}
}
{{end}}
