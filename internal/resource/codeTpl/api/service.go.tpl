package svc{{.PackagePascalName}}

import (
    "{{.ImportDirPrefix}}/demo/dto/dtoUser"

    "github.com/gin-gonic/gin"
)

{{if .TargetFileExist}}
type {{.ReceiverTypePascalName}}Svc interface {
    {{.FunctionName}}(c *gin.Context)
}

type {{.ReceiverTypeName}}Svc struct {
}

var _ {{.ReceiverTypePascalName}}Svc = (*{{.ReceiverTypeName}}Svc)(nil)

func New{{.ReceiverTypePascalName}}Svc() {{.ReceiverTypePascalName}}Svc {
    return &{{.ReceiverTypeName}}Svc{
    }
}
{{end}}

func (svc *{{.ReceiverTypeName}}Svc) {{.FunctionName}}(c *gin.Context, req *dto{{.PackagePascalName}}.{{.FunctionName}}Req) (*dto{{.PackagePascalName}}.{{.FunctionName}}Res, error) {
    return &dto{{.PackagePascalName}}.{{.FunctionName}}Res{}, nil
}
