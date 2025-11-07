package dtouser

import (
	"github.com/morehao/go-gin-web/apps/demoapp/object/objcommon"
	"github.com/morehao/go-gin-web/apps/demoapp/object/objuser"
)

type UserCreateReq struct {
	objuser.UserBaseInfo
}

type UserUpdateReq struct {
	ID uint `json:"id" validate:"required" label:"数据自增id"` // 数据自增id
	objuser.UserBaseInfo
}

type UserDetailReq struct {
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}

type UserPageListReq struct {
	objcommon.PageQuery
}

type UserDeleteReq struct {
	ID uint `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}
