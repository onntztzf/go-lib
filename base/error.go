package base

type Error struct {
	Code int
	Msg  string
}

func NewError(code int, msg string) Error {
	return Error{Code: code, Msg: msg}
}

func NewErrorWithMsg(msg string) Error {
	return Error{Code: ErrorCodeCustomError, Msg: msg}
}

func (err Error) Error() string {
	return err.Msg
}

func (err Error) ReplaceMsg(newValue string) Error {
	err.Msg = newValue
	return err
}

const (
	ErrorCodeCustomError = iota + 1
	ErrorCodeSystemError
	ErrorCodeDBInsert
	ErrorCodeDBSelect
	ErrorCodeDBUpdate
	ErrorCodeDBDelete
	ErrorCodeInvalidParam

	ErrorCodeNotLogin = iota + 200
)

var ErrorSystemError = NewError(ErrorCodeSystemError, "系统错误")
var ErrorInvalidParam = NewError(ErrorCodeInvalidParam, "参数错误")

var ErrorDBInsert = NewError(ErrorCodeDBInsert, "添加记录失败")
var ErrorDBSelect = NewError(ErrorCodeDBSelect, "查询记录失败")
var ErrorDBUpdate = NewError(ErrorCodeDBUpdate, "更新记录失败")
var ErrorDBDelete = NewError(ErrorCodeDBDelete, "删除记录失败")

var ErrorNotLogin = NewError(ErrorCodeNotLogin, "用户未登录")
