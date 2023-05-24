package e

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

const (
	ErrCodeSystemError        = 1
	ErrCodeCustomError        = 2
	ErrCodeInvalidParam       = 3
	ErrCodeDatabaseInsertFail = 4
	ErrCodeDatabaseSelectFail = 5
	ErrCodeDatabaseUpdateFail = 6
	ErrCodeDatabaseDeleteFail = 7
	ErrCodeDatabaseCommitFail = 8
	ErrCodeNoData             = 9
)

var SystemError = NewError(ErrCodeSystemError, "system error")

var ErrInvalidParam = NewError(ErrCodeInvalidParam, "invalid param")
var ErrDatabaseInsertFail = NewError(ErrCodeDatabaseInsertFail, "failed to add record")
var ErrDatabaseSelectFail = NewError(ErrCodeDatabaseSelectFail, "failed to retrieve record")
var ErrDatabaseUpdateFail = NewError(ErrCodeDatabaseUpdateFail, "failed to update record")
var ErrDatabaseDeleteFail = NewError(ErrCodeDatabaseDeleteFail, "failed to delete record")
var ErrDatabaseCommitFail = NewError(ErrCodeDatabaseCommitFail, "failed to commit")

var ErrNoData = NewError(ErrCodeNoData, "no data")
