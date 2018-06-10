package middlewares

import (
	"err"
	"model"
	"response"

	"github.com/kataras/iris"
)

// CheckPostIDExistence checks the existence of the post.
func CheckPostIDExistence(ctx iris.Context) {
	postID, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)

	has, er := model.CheckPostByID(postID)
	err.CheckErrWithPanic(er)
	if !has {
		response.NotFound(ctx, iris.Map{})
		ctx.StopExecution()
		return
	}

	ctx.Next()
}

// CheckCommentIDExistenceAndLegitimate checks whether the post and the comment exists,
// together with whether the comment belongs to the post.
func CheckCommentIDExistenceAndLegitimate(ctx iris.Context) {
	postID, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)

	commentID, er := ctx.Params().GetInt64("commentId")
	err.CheckErrWithPanic(er)

	has, er := model.CheckCommentByID(commentID)
	err.CheckErrWithPanic(er)
	if !has {
		response.NotFound(ctx, iris.Map{})
		ctx.StopExecution()
		return
	}

	has, er = model.CheckCommentByPost(postID, commentID)
	err.CheckErrWithPanic(er)
	if !has {
		response.Forbidden(ctx, iris.Map{})
		ctx.StopExecution()
		return
	}

	ctx.Next()
}
