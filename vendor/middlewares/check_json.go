package middlewares

import (
	"response"

	"github.com/kataras/iris"
)

const (
	jsonContentType string = "application/json"
)

// CheckContentTypeJSON ...
func CheckContentTypeJSON(ctx iris.Context) {
	if ctx.Method() != iris.MethodGet &&
		ctx.Method() != iris.MethodOptions &&
		ctx.Method() != iris.MethodDelete {
		requestContentType := ctx.GetHeader("Content-Type")
		if requestContentType != jsonContentType {
			response.BadRequest(ctx, iris.Map{})
			ctx.StopExecution()
		}
	}

	ctx.Next()
}

// SetResponseContentTypeJSON ...
func SetResponseContentTypeJSON(ctx iris.Context) {
	ctx.ContentType(jsonContentType)

	ctx.Next()
}
