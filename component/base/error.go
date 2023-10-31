package base

import (
	"fmt"
	"github.com/pkg/errors"
)

type Error struct {
	Code int
	Msg  string
}

func NewBaseError(code int, message string) *Error {
	return &Error{
		Code: code,
		Msg:  message,
	}
}

func NewError(code int, message, userMsg string) Error {
	return Error{
		Code: code,
		Msg:  message,
	}
}

func (err Error) Error() string {
	return err.Msg
}

func (err Error) Sprintf(v ...interface{}) Error {
	err.Msg = fmt.Sprintf(err.Msg, v...)
	return err
}

func (err Error) Equal(e error) bool {
	switch errors.Cause(e).(type) {
	case Error:
		return err.Code == errors.Cause(e).(Error).Code
	default:
		return false
	}
}

func (err Error) WrapPrint(core error, message string, user ...interface{}) error {
	if core == nil {
		return nil
	}
	err.SetErrPrintfMsg(core)
	return errors.Wrap(err, message)
}

func (err Error) WrapPrintf(core error, format string, message ...interface{}) error {
	if core == nil {
		return nil
	}
	err.SetErrPrintfMsg(core)
	return errors.Wrap(err, fmt.Sprintf(format, message...))
}

func (err Error) Wrap(core error) error {
	if core == nil {
		return nil
	}

	msg := err.Msg
	err.Msg = core.Error()
	return errors.Wrap(err, msg)
}

func (err *Error) SetErrPrintfMsg(v ...interface{}) {
	err.Msg = fmt.Sprintf(err.Msg, v...)
}

// deprecated: use Error.SetErrPrintfMsg instead
func SetErrPrintfMsg(err *Error, v ...interface{}) {
	err.Msg = fmt.Sprintf(err.Msg, v...)
}
