package dtoUser

import (
	"go-gin-web/internal/app/object/objCommon"
	"go-gin-web/internal/app/object/objUser"
)

type UserCreateResp struct {
	Id uint64 `json:"id"` // 数据自增id
}

type UserDetailResp struct {
	Id uint64 `json:"id" validate:"required"` // 数据自增id
	objUser.UserBaseInfo
	objCommon.OperatorBaseInfo
}

type UserPageListItem struct {
	Id uint64 `json:"id" validate:"required"` // 数据自增id
	objUser.UserBaseInfo
	objCommon.OperatorBaseInfo
}

type UserPageListResp struct {
	List  []UserPageListItem `json:"list"`  // 数据列表
	Total int64              `json:"total"` // 数据总条数
}
