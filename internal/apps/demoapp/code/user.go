package code

import "github.com/morehao/golib/gerror"

const (
	UserCreateErr      = 100100
	UserDeleteErr      = 100101
	UserUpdateErr      = 100102
	UserGetDetailErr   = 100103
	UserGetPageListErr = 100104
	UserNotExistErr    = 100105
)

var userErrorMsgMap = gerror.CodeMsgMap{
	UserCreateErr:      "创建用户失败",
	UserDeleteErr:      "删除用户失败",
	UserUpdateErr:      "修改用户失败",
	UserGetDetailErr:   "查看用户失败",
	UserGetPageListErr: "查看用户列表失败",
	UserNotExistErr:    "用户不存在",
}
