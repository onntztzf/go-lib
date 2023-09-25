package gin_output

import (
	"net/http"

	"github.com/2hangpeng/go-lib/e"
	"github.com/gin-gonic/gin"
)

type Output struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Success sends a success response with optional data.
func Success(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	ctx.JSON(http.StatusOK, Output{
		Code: 0,
		Msg:  "success",
		Data: data,
	})
}

// Failure sends a failure response with an error.
func Failure(ctx *gin.Context, err error) {
	code, msg := e.SystemError.Code, e.SystemError.Msg
	if customErr, ok := err.(e.Error); ok {
		code = customErr.Code
		msg = customErr.Msg
	}
	ctx.JSON(http.StatusOK, Output{
		Code: code,
		Msg:  msg,
		Data: struct{}{},
	})
}
