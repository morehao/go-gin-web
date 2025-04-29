package dtoUser

import (
	"go-gin-web/apps/demo/internal/object/objCommon"
	"go-gin-web/apps/demo/internal/object/objUser"
)

type UserCreateReq struct {
	objUser.UserBaseInfo
}

type UserUpdateReq struct {
	ID uint64 `json:"id" validate:"required" label:"数据自增id"` // 数据自增id
	objUser.UserBaseInfo
}

type UserDetailReq struct {
	ID uint64 `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}

type UserPageListReq struct {
	objCommon.PageQuery
}

type UserDeleteReq struct {
	ID uint64 `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}
