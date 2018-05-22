package middlewares

import (
	"response"

	"github.com/kataras/iris"
)

// InternalErrorCatcher ..
func InternalErrorCatcher(ctx iris.Context) {
	defer func() {
		if rval := recover(); rval != nil {
			response.InternalServerError(ctx, iris.Map{})
			ctx.StopExecution()
		}
	}()

	ctx.Next()
}
