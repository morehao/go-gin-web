package dto{{.PackagePascalName}}

type {{.StructName}}CreateResp struct {
	Id uint64 `json:"id"` // 数据自增id
}

type {{.StructName}}DetailResp struct {
	Id        uint64 `json:"id" validate:"required"` // 数据自增id
	CreatedAt int64  `json:"createdAt"`              // 系统创建时间
}

type {{.StructName}}PageListItem struct {
	Id        uint64 `json:"id" validate:"required"` // 数据自增id
	CreatedAt int64  `json:"createdAt"`              // 系统创建时间
}

type {{.StructName}}PageListResp struct {
	List  []{{.StructName}}PageListItem `json:"list"`  // 数据列表
	Total int64          `json:"total"` // 数据总条数
}
