package code

import "github.com/morehao/golib/gerror"

const (
	UserCreateError      = 100100
	UserDeleteError      = 100101
	UserUpdateError      = 100102
	UserGetDetailError   = 100103
	UserGetPageListError = 100104
	UserNotExistError    = 100105
)

var userErrorMsgMap = gerror.CodeMsgMap{
	UserCreateError:      "创建用户管理失败",
	UserDeleteError:      "删除用户管理失败",
	UserUpdateError:      "修改用户管理失败",
	UserGetDetailError:   "查看用户管理失败",
	UserGetPageListError: "查看用户管理列表失败",
	UserNotExistError:    "用户管理不存在",
}
