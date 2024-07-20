package ctr{{.PackagePascalName}}

import (
	"{{.ImportDirPrefix}}/dto/dto{{.PackagePascalName}}"
	"{{.ImportDirPrefix}}/service/svc{{.PackagePascalName}}"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/gcontext/ginRender"
)

{{if .TargetFileExist}}
type {{.ReceiverTypePascalName}}Ctr interface {
	Create(c *gin.Context)
}

type {{.ReceiverTypeName}}Ctr struct {
	{{.ReceiverTypeName}}Svc svc{{.ReceiverTypePascalName}}.{{.ReceiverTypePascalName}}Svc
}

var _ {{.ReceiverTypePascalName}}Ctr = (*{{.ReceiverTypeName}}Ctr)(nil)

func New{{.ReceiverTypePascalName}}Ctr() {{.ReceiverTypePascalName}}Ctr {
	return &{{.ReceiverTypeName}}Ctr{
		{{.ReceiverTypePascalName}}Svc: svc{{.ReceiverTypePascalName}}.New{{.ReceiverTypePascalName}}Svc(),
	}
}

{{end}}

{{if eq .HttpMethod "POST"}}
// Create {{.Description}}
// @Tags {{.ApiDocTag}}
// @Summary {{.Description}}
// @accept application/json
// @Produce application/json
// @Param req body dto{{.PackagePascalName}}.CreateReq true "{{.Description}}"
// @Success 200 {object} dto{{.PackagePascalName}}.CreateRes "{"code": 0,"data": "ok","msg": "success"}"
// @Router {{.ApiPath}} [post]
func (ctr *{{.ReceiverTypeName}}Ctr) Create(c *gin.Context) {
	var req dto{{.PackagePascalName}}.CreateReq
	if err := c.ShouldBindJSON(c, &req); err != nil {
		base.RenderJsonFail(c, err)
		return
	}
	res, err := ctr.{{.ReceiverTypeName}}Svc.Create(c, &req)
	if err != nil {
		base.RenderJsonFail(c, err)
		return
	} else {
		base.RenderJsonSucc(c, res)
	}
}
{{else if eq .HttpMethod "GET"}}
// Create {{.Description}}
// @Tags {{.ApiDocTag}}
// @Summary {{.Description}}
// @accept application/json
// @Produce application/json
// @Param req query dto{{.PackagePascalName}}.CreateReq true "{{.Description}}"
// @Success 200 {object} dto{{.PackagePascalName}}.CreateRes "{"code": 0,"data": "ok","msg": "success"}"
// @Router {{.ApiPath}} [get]
func (ctr *{{.ReceiverTypeName}}Ctr)Create(c *gin.Context) {
	var req dto{{.PackagePascalName}}.CreateReq
	if err := c.ShouldBindQuery(c, &req); err != nil {
		base.RenderJsonFail(c, err)
		return
	}
	res, err := ctr.{{.ReceiverTypeName}}Svc.Create(c, &req)
	if err != nil {
		base.RenderJsonFail(c, err)
		return
	} else {
		base.RenderJsonSucc(c, res)
	}
}
{{end}}
