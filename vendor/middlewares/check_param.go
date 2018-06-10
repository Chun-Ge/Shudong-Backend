package middlewares

import (
	"err"
	"model"
	"response"

	"github.com/kataras/iris"
)

// CheckPostIDExistence .
func CheckPostIDExistence(ctx iris.Context) {
	postID, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)

	has, er := model.CheckPostByPostID(postID)
	err.CheckErrWithPanic(er)
	if !has {
		response.NotFound(ctx, iris.Map{})
		ctx.StopExecution()
		return
	}

	ctx.Next()
}
