package ctr{{.PackagePascalName}}

import (
	"{{.ImportDirPrefix}}/dto/dto{{.PackagePascalName}}"
	"{{.ImportDirPrefix}}/service/svc{{.PackagePascalName}}"

	"github.com/gin-gonic/gin"
	"github.com/morehao/go-tools/gcontext/ginRender"
)

type {{.ReceiverTypePascalName}}Ctr interface {
	Create(c *gin.Context)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	Detail(c *gin.Context)
	PageList(c *gin.Context)
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


// Create {{.Description}}
// @Tags {{.ApiDocTag}}
// @Summary {{.Description}}
// @accept application/json
// @Produce application/json
// @Param req body dto{{.PackagePascalName}}.{{.StructName}}CreateReq true "{{.Description}}"
// @Success 200 {object} dto{{.PackagePascalName}}.{{.StructName}}CreateResp "{"code": 0,"data": "ok","msg": "success"}"
// @Router {{.ModuleApiPrefix}}/create [post]
func (ctr *{{.ReceiverTypeName}}Ctr) Create(c *gin.Context) {
	var req dto{{.PackagePascalName}}.{{.StructName}}CreateReq
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

// Delete {{.Description}}
// @Tags {{.ApiDocTag}}
// @Summary {{.Description}}
// @accept application/json
// @Produce application/json
// @Param req body dto{{.PackagePascalName}}.{{.StructName}}DeleteReq true "{{.Description}}"
// @Router {{.ModuleApiPrefix}}/delete [post]
func (ctr *{{.ReceiverTypeName}}Ctr) Delete(c *gin.Context) {
	var req dto{{.PackagePascalName}}.{{.StructName}}DeleteReq
	if err := c.ShouldBindJSON(c, &req); err != nil {
		base.RenderJsonFail(c, err)
		return
	}

	if err := ctr.{{.ReceiverTypeName}}Svc.Delete(c, &req); err != nil {
		base.RenderJsonFail(c, err)
		return
	} else {
		base.RenderJsonSucc(c, "删除成功")
	}
}

// Update {{.Description}}
// @Tags {{.ApiDocTag}}
// @Summary {{.Description}}
// @accept application/json
// @Produce application/json
// @Param req body dto{{.PackagePascalName}}.{{.StructName}}UpdateReq true "{{.Description}}"
// @Router {{.ModuleApiPrefix}}/update [post]
func (ctr *{{.ReceiverTypeName}}Ctr) Update(c *gin.Context) {
	var req dto{{.PackagePascalName}}.{{.StructName}}UpdateReq
	if err := c.ShouldBindJSON(c, &req); err != nil {
		base.RenderJsonFail(c, err)
		return
	}
	if err := ctr.{{.ReceiverTypeName}}Svc.Update(c, &req); err != nil {
		base.RenderJsonFail(c, err)
		return
	} else {
		base.RenderJsonSucc(c, "修改成功")
	}
}

// Detail {{.Description}}
// @Tags {{.ApiDocTag}}
// @Summary {{.Description}}
// @accept application/json
// @Produce application/json
// @Param req query dto{{.PackagePascalName}}.{{.StructName}}DetailReq true "{{.Description}}"
// @Success 200 {object} dto{{.PackagePascalName}}.{{.StructName}}DetailResp "{"code": 0,"data": "ok","msg": "success"}"
// @Router {{.ModuleApiPrefix}}/detail [get]
func (ctr *{{.ReceiverTypeName}}Ctr) Detail(c *gin.Context) {
	var req dto{{.PackagePascalName}}.{{.StructName}}DetailReq
	if err := c.ShouldBindQuery(c, &req); err != nil {
		base.RenderJsonFail(c, err)
		return
	}
	res, err := ctr.{{.ReceiverTypeName}}Svc.Detail(c, &req)
	if err != nil {
		base.RenderJsonFail(c, err)
		return
	} else {
		base.RenderJsonSucc(c, res)
	}
}

// PageList {{.Description}}
// @Tags {{.ApiDocTag}}
// @Summary {{.Description}}
// @accept application/json
// @Produce application/json
// @Param req query dto{{.PackagePascalName}}.{{.StructName}}PageListReq true "{{.Description}}"
// @Success 200 {object} dto{{.PackagePascalName}}.{{.StructName}}PageListResp "{"code": 0,"data": "ok","msg": "success"}"
// @Router {{.ModuleApiPrefix}}/pageList [get]
func (ctr *{{.ReceiverTypeName}}Ctr) PageList(c *gin.Context) {
	var req dto{{.PackagePascalName}}.{{.StructName}}PageListReq
	if err := c.ShouldBindQuery(c, &req); err != nil {
		base.RenderJsonFail(c, err)
		return
	}
	res, err := ctr.{{.ReceiverTypeName}}Svc.PageList(c, &req)
	if err != nil {
		base.RenderJsonFail(c, err)
		return
	} else {
		base.RenderJsonSucc(c, res)
	}
}
