package service

import (
	"err"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// delete a post
// route: /post/{postid:int min(1)}
// pre: the post belongs to the user
// post: the post has been deleted, meanwhile,
//       all comments of the post will have been deleted,
//       and clear the info of upvoting the post
func DeletePost(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postid")
	err.CheckErrWithPanic(er)

	has, er := model.CheckPostByUser(userID, postID)
	err.CheckErrWithPanic(er)

	// if the post do not belong to the user
	if has == false {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	// delete the upvoting info of the post
	_, er = model.CancelUpvotePostByPost(postID)
	err.CheckErrWithPanic(er)

	// delete all comments of the post
	_, er = model.CancelCommentByPost(postID)
	err.CheckErrWithPanic(er)

	// delete the post
	_, er = model.CancelPostByID(postID)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{})
}
