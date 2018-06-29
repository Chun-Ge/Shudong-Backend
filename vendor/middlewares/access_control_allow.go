package middlewares

import (
	"github.com/kataras/iris"
)

// SetAccessControlAllowOrigin ...
func SetAccessControlAllowOrigin(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")

	ctx.Next()
}
