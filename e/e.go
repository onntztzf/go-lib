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

var SystemError = NewError(SystemErrorCode, "系统错误")
var InvalidParam = NewError(InvalidParamErrorCode, "参数错误")

var DatabaseInsertError = NewError(DatabaseInsertErrorCode, "添加记录失败")
var DatabaseSelectError = NewError(DatabaseSelectErrorCode, "查询记录失败")
var DatabaseUpdateError = NewError(DatabaseUpdateErrorCode, "更新记录失败")
