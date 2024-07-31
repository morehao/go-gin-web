package errorCode

import "github.com/morehao/go-tools/gerror"

var UserCreateErr = gerror.Error{
	Code: 100100,
	Msg:  "创建用户失败",
}

var UserDeleteErr = gerror.Error{
	Code: 100101,
	Msg:  "删除用户失败",
}

var UserUpdateErr = gerror.Error{
	Code: 100102,
	Msg:  "修改用户失败",
}

var UserGetDetailErr = gerror.Error{
	Code: 100103,
	Msg:  "查看用户失败",
}

var UserGetPageListErr = gerror.Error{
	Code: 100104,
	Msg:  "查看用户列表失败",
}

var UserNotExistErr = gerror.Error{
	Code: 100105,
	Msg:  "用户已存在",
}
