package e

import (
	"fmt"
)

type Error struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) Error {
	return Error{Code: code, Msg: msg}
}

func NewErrorWithMsg(msg string) Error {
	return Error{Code: ErrCodeCustomError, Msg: msg}
}

func (err Error) Error() string {
	return err.Msg
}

func (err Error) ReplaceMsg(newValue string) Error {
	err.Msg = newValue
	return err
}

func (err Error) ReplaceMsgf(format string, args ...interface{}) Error {
	err.Msg = fmt.Sprintf(format, args...)
	return err
}

const (
	ErrCodeSystemError        = 1
	ErrCodeCustomError        = 2
	ErrCodeInvalidParam       = 3
	ErrCodeDatabaseInsertFail = 4
	ErrCodeDatabaseSelectFail = 5
	ErrCodeDatabaseUpdateFail = 6
	ErrCodeDatabaseDeleteFail = 7
	ErrCodeDatabaseCommitFail = 8
)

var SystemError = NewError(ErrCodeSystemError, "system error")
var InvalidParam = NewError(ErrCodeInvalidParam, "invalid param")
var DatabaseInsertFail = NewError(ErrCodeDatabaseInsertFail, "failed to add")
var DatabaseSelectFail = NewError(ErrCodeDatabaseSelectFail, "failed to select")
var DatabaseUpdateFail = NewError(ErrCodeDatabaseUpdateFail, "failed to update")
var DatabaseDeleteFail = NewError(ErrCodeDatabaseDeleteFail, "failed to delete")
var DatabaseCommitFail = NewError(ErrCodeDatabaseCommitFail, "failed to commit")
