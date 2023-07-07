package gin_output

import (
	"net/http"

	"github.com/2hangpeng/go-lib/e"
	"github.com/gin-gonic/gin"
)

type output struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(ctx *gin.Context, data interface{}) {
	if data == nil {
		data = struct{}{}
	}
	ctx.JSON(http.StatusOK, output{
		Code: 0,
		Msg:  "",
		Data: data,
	})
}

func Failure(ctx *gin.Context, err error) {
	code, msg := e.SystemError.Code, e.SystemError.Msg
	if err, ok := err.(e.Error); ok {
		code = err.Code
		msg = err.Msg
	}
	ctx.JSON(http.StatusOK, output{
		Code: code,
		Msg:  msg,
		Data: struct{}{},
	})
}
