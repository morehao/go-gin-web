package dto{{.PackagePascalName}}

import "{{.ImportDirPrefix}}/object/objCommon"

type {{.StructName}}CreateReq struct {
}

type {{.StructName}}UpdateReq struct {
	Id   uint64 `json:"id" validate:"required" label:"数据自增id"` // 数据自增id
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
