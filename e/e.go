package e

type Error struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) Error {
	return Error{Code: code, Msg: msg}
}

func NewErrorWithMsg(msg string) Error {
	return Error{Code: CustomErrorCode, Msg: msg}
}

func (err Error) Error() string {
	return err.Msg
}

func (err Error) ReplaceMsg(newValue string) Error {
	err.Msg = newValue
	return err
}

const (
	SystemErrorCode         = 1
	CustomErrorCode         = 2
	InvalidParamErrorCode   = 3
	DatabaseInsertErrorCode = 4
	DatabaseSelectErrorCode = 5
	DatabaseUpdateErrorCode = 6
)

var SystemError = NewError(SystemErrorCode, "system error")
var InvalidParam = NewError(InvalidParamErrorCode, "invalid param")

var DatabaseInsertError = NewError(DatabaseInsertErrorCode, "failed to add record")
var DatabaseSelectError = NewError(DatabaseSelectErrorCode, "failed to retrieve record")
var DatabaseUpdateError = NewError(DatabaseUpdateErrorCode, "failed to update record")
