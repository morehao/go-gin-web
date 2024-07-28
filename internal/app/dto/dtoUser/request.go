package dtoUser

import (
	"go-gin-web/internal/app/object/objCommon"
	"go-gin-web/internal/app/object/objUser"
)

type UserCreateReq struct {
	objUser.UserBaseInfo
}

type UserUpdateReq struct {
	Id uint64 `json:"id" validate:"required" label:"数据自增id"` // 数据自增id
	objUser.UserBaseInfo
}

type UserDetailReq struct {
	Id uint64 `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}

type UserPageListReq struct {
	objCommon.PageQuery
}

type UserDeleteReq struct {
	Id uint64 `json:"id" form:"id" validate:"required" label:"数据自增id"` // 数据自增id
}
