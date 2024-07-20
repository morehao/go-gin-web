package router

import (
	"{{.ImportDirPrefix}}/controller/ctr{{.PackagePascalName}}"

	"github.com/gin-gonic/gin"
)

// {{.ReceiverTypeName}}Router 初始化 {{.Description}} 路由信息
func {{.ReceiverTypeName}}Router(privateRouter *gin.RouterGroup) {
	{{.ReceiverTypeName}}Ctr := ctr{{.PackagePascalName}}.New{{.ReceiverTypePascalName}}Ctr()
	routerGroup := privateRouter.Group("{{.ModuleApiPrefix}}")
	{
		routerGroup.POST("create", {{.ReceiverTypeName}}Ctr.Create)   // 新建{{.Description}}
		routerGroup.POST("delete", {{.ReceiverTypeName}}Ctr.Delete)   // 删除{{.Description}}
		routerGroup.POST("update", {{.ReceiverTypeName}}Ctr.Update)   // 更新{{.Description}}
		routerGroup.GET("detail", {{.ReceiverTypeName}}Ctr.Detail)    // 根据ID获取{{.Description}}
        routerGroup.GET("pageList", {{.ReceiverTypeName}}Ctr.PageList)  // 获取{{.Description}}列表
	}
}
