package ctr{{.PackagePascalName}}

import (
	"{{.ImportDirPrefix}}/dto/dto{{.PackagePascalName}}"
	"{{.ImportDirPrefix}}/service/svc{{.PackagePascalName}}"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/gcontext/ginRender"
)
{{if not .TargetFileExist}}
type {{.ReceiverTypePascalName}}Ctr interface {
	{{.FunctionName}}(c *gin.Context)
}

type {{.ReceiverTypeName}}Ctr struct {
	{{.ReceiverTypeName}}Svc svc{{.ReceiverTypePascalName}}.{{.ReceiverTypePascalName}}Svc
}

var _ {{.ReceiverTypePascalName}}Ctr = (*{{.ReceiverTypeName}}Ctr)(nil)

func New{{.ReceiverTypePascalName}}Ctr() {{.ReceiverTypePascalName}}Ctr {
	return &{{.ReceiverTypeName}}Ctr{
		{{.ReceiverTypeName}}Svc: svc{{.ReceiverTypePascalName}}.New{{.ReceiverTypePascalName}}Svc(),
	}
}
{{end}}
{{if eq .HttpMethod "POST"}}
// {{.FunctionName}} {{.Description}}
// @Tags {{.ApiDocTag}}
// @Summary {{.Description}}
// @accept application/json
// @Produce application/json
// @Param req body dto{{.PackagePascalName}}.{{.FunctionName}}Req true "{{.Description}}"
// @Success 200 {object} dto{{.PackagePascalName}}.{{.FunctionName}}Resp "{"code": 0,"data": "ok","msg": "success"}"
// @Router {{.ApiPath}} [post]
func (ctr *{{.ReceiverTypeName}}Ctr) {{.FunctionName}}(c *gin.Context) {
	var req dto{{.PackagePascalName}}.{{.FunctionName}}Req
	if err := c.ShouldBindJSON(&req); err != nil {
		ginRender.Fail(c, err)
		return
	}
	res, err := ctr.{{.ReceiverTypeName}}Svc.{{.FunctionName}}(c, &req)
	if err != nil {
		ginRender.Fail(c, err)
		return
	} else {
		ginRender.Success(c, res)
	}
}
{{else if eq .HttpMethod "GET"}}
// Create {{.Description}}
// @Tags {{.ApiDocTag}}
// @Summary {{.Description}}
// @accept application/json
// @Produce application/json
// @Param req query dto{{.PackagePascalName}}.{{.FunctionName}}Req true "{{.Description}}"
// @Success 200 {object} dto{{.PackagePascalName}}.{{.FunctionName}}Resp "{"code": 0,"data": "ok","msg": "success"}"
// @Router {{.ApiPath}} [get]
func (ctr *{{.ReceiverTypeName}}Ctr){{.FunctionName}}(c *gin.Context) {
	var req dto{{.PackagePascalName}}.{{.FunctionName}}Req
	if err := c.ShouldBindQuery(&req); err != nil {
		ginRender.Fail(c, err)
		return
	}
	res, err := ctr.{{.ReceiverTypeName}}Svc.{{.FunctionName}}(c, &req)
	if err != nil {
		ginRender.Fail(c, err)
		return
	} else {
		ginRender.Success(c, res)
	}
}
{{end}}
