package base

type Error struct {
	Code int
	Msg  string
}

func GetError(code int, msg string) Error {
	return Error{Code: code, Msg: msg}
}

func GetErrorWithMsg(msg string) Error {
	return Error{Code: ErrorCodeCustomError, Msg: msg}
}

func (err Error) Error() string {
	return err.Msg
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

var ErrorSystemError = GetError(ErrorCodeSystemError, "系统错误")
var ErrorInvalidParam = GetError(ErrorCodeInvalidParam, "参数错误")

var ErrorDBInsert = GetError(ErrorCodeDBInsert, "添加记录失败")
var ErrorDBSelect = GetError(ErrorCodeDBSelect, "查询记录失败")
var ErrorDBUpdate = GetError(ErrorCodeDBUpdate, "更新记录失败")
var ErrorDBDelete = GetError(ErrorCodeDBDelete, "删除记录失败")

var ErrorNotLogin = GetError(ErrorCodeNotLogin, "用户未登录")
