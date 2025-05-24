package code

import "github.com/morehao/golib/gerror"

const (
	DBInsertErr = 100000
	DBDeleteErr = 100001
	DBUpdateErr = 100002
	DBFindErr   = 100003
)

var dbErrorMsgMap = gerror.CodeMsgMap{
	DBInsertErr: "db insert error",
	DBDeleteErr: "db delete error",
	DBUpdateErr: "db update error",
	DBFindErr:   "db find error",
}
