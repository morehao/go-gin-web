package dtoUser

import (
	"go-gin-web/apps/demo/internal/object/objCommon"
	"go-gin-web/apps/demo/internal/object/objUser"
)

type UserCreateResp struct {
	ID uint64 `json:"id"` // 数据自增id
}

type UserDetailResp struct {
	ID uint64 `json:"id" validate:"required"` // 数据自增id
	objUser.UserBaseInfo
	objCommon.OperatorBaseInfo
}

type UserPageListItem struct {
	ID uint64 `json:"id" validate:"required"` // 数据自增id
	objUser.UserBaseInfo
	objCommon.OperatorBaseInfo
}

type UserPageListResp struct {
	List  []UserPageListItem `json:"list"`  // 数据列表
	Total int64              `json:"total"` // 数据总条数
}
