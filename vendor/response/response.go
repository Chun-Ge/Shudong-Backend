package response

import (
	"net/http"

	"github.com/kataras/iris"
)

// OuterMsg : sturct of the response msg
type OuterMsg struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func genResponseMsg(messageCode int, data interface{}) *OuterMsg {
	message := http.StatusText(messageCode)
	return &OuterMsg{
		Msg:  message,
		Data: data,
	}
}

// OK ...
func OK(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	ctx.JSON(genResponseMsg(http.StatusOK, data))
}

// Created ...
func Created(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(genResponseMsg(http.StatusCreated, data))
}

// BadRequest ...
func BadRequest(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusBadRequest)
	ctx.JSON(genResponseMsg(http.StatusBadRequest, data))
}

// Unauthorized ...
func Unauthorized(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusUnauthorized)
	ctx.JSON(genResponseMsg(http.StatusUnauthorized, data))
}

// Forbidden ...
func Forbidden(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusForbidden)
	ctx.JSON(genResponseMsg(http.StatusForbidden, data))
}

// NotFound ...
func NotFound(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusNotFound)
	ctx.JSON(genResponseMsg(http.StatusNotFound, data))
}

// Conflict ...
func Conflict(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusConflict)
	ctx.JSON(genResponseMsg(http.StatusConflict, data))
}

// InternalServerError ...
func InternalServerError(ctx iris.Context, data interface{}) {
	ctx.StatusCode(iris.StatusInternalServerError)
	ctx.JSON(genResponseMsg(http.StatusInternalServerError, data))
}
