package routerHttp

import (
	"{{.ImportDirPrefix}}/ctr{{.PackagePascalName}}"

	"github.com/gin-gonic/gin"
)

// {{.PackageName}}Router 初始化 {{.Description}} 路由信息
func {{.PackageName}}Router(privateRouter *gin.RouterGroup) {
	{{.PackageName}}Ctr := ctr{{.PackagePascalName}}.New{{.PackagePascalName}}Ctr()
	routerGroup := privateRouter.Group("{{.PackageName}}")
	{
		routerGroup.POST("create", ctr{{.PackagePascalName}}.Create)   // 新建{{.Description}}
		routerGroup.POST("delete", ctr{{.PackagePascalName}}.Delete)   // 删除{{.Description}}
		routerGroup.POST("update", ctr{{.PackagePascalName}}.Update)   // 更新{{.Description}}
		routerGroup.GET("detail", ctr{{.PackagePascalName}}.Detail)    // 根据ID获取{{.Description}}
        routerGroup.GET("pageList", ctr{{.PackagePascalName}}.PageList)  // 获取{{.Description}}列表
	}
}
