package errorCode

import "github.com/morehao/go-tools/gerror"

var {{.StructName}}CreateErr = gerror.Error{
	Code: 100100,
	Msg:  "创建{{.Description}}失败",
}

var {{.StructName}}DeleteErr = gerror.Error{
	Code: 100101,
	Msg:  "删除{{.Description}}失败",
}

var {{.StructName}}UpdateErr = gerror.Error{
	Code: 100102,
	Msg:  "修改{{.Description}}失败",
}

var Get{{.StructName}}DetailErr = gerror.Error{
	Code: 100103,
	Msg:  "查看{{.Description}}失败",
}

var Get{{.StructName}}PageListErr = gerror.Error{
	Code: 100100,
	Msg:  "查看{{.Description}}列表失败",
}
