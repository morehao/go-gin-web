package router

import (
	"{{.ProjectRootDir}}/internal/{{.ServiceName}}/controller/ctr{{.PackagePascalName}}"

	"github.com/gin-gonic/gin"
)
{{if not .TargetFileExist}}
// {{.ReceiverTypeName}}Router 初始化{{.Description}}路由信息
func {{.ReceiverTypeName}}Router(routerGroup *gin.RouterGroup) {
	{{.ReceiverTypeName}}Ctr := ctr{{.PackagePascalName}}.New{{.ReceiverTypePascalName}}Ctr()
	{{.ReceiverTypeName}}Group := routerGroup.Group("{{.ApiGroup}}")
	{
		{{.ReceiverTypeName}}Group.{{.HttpMethod}}("{{.ApiSuffix}}", {{.ReceiverTypeName}}Ctr.{{.FunctionName}})   // {{.Description}}
	}
}
{{end}}
