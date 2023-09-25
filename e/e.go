package e

// Error is a custom error type with code and message.
type Error struct {
	Code int
	Msg  string
}

const (
	ErrCodeSystemError   = 1
	ErrCodeInvalidParam  = 2
	ErrCodeDatabaseError = 3
)

var (
	SystemError   = NewError(ErrCodeSystemError, "system error")
	InvalidParam  = NewError(ErrCodeInvalidParam, "invalid param")
	DatabaseError = NewError(ErrCodeDatabaseError, "database error")
)

// NewError creates a new custom error with the given code and message.
func NewError(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

// Error returns the error message.
func (err Error) Error() string {
	return err.Msg
}
