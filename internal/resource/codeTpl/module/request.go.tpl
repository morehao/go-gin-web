package dto{{.PackagePascalName}}

import (
	"{{.ProjectRootDir}}/internal/app/object/objCommon"
	"{{.ProjectRootDir}}/internal/app/object/obj{{.PackagePascalName}}"
)

type {{.StructName}}CreateReq struct {
	obj{{.PackagePascalName}}.{{.StructName}}BaseInfo
}

type {{.StructName}}UpdateReq struct {
	Id uint64 `json:"id" validate:"required" label:"数据自增id"` // 数据自增id
	obj{{.PackagePascalName}}.{{.StructName}}BaseInfo
}

type {{.StructName}}DetailReq struct {
	Id uint64 `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}

type {{.StructName}}PageListReq struct {
	objCommon.PageQuery
}

type {{.StructName}}DeleteReq struct {
	Id uint64 `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}
