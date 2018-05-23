package service

import (
	"err"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// delete comment
// route: /post/{postid}/comments/{commentid}
// pre: the comment belongs to the post
// post: the comment has been deleted, meanwhile,
//       the upvoting info of the comment will have been deleted
// response result: 1. OK
//                  2. Forbidden: post do not belongs to the user
//                     or comment do not belongs to the post
//                  3. Unauthorized: the user is not valid
func DeleteComment(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postid")
	err.CheckErrWithPanic(er)
	commentID, er := ctx.Params().GetInt64("commentid")
	err.CheckErrWithPanic(er)

	has, er := model.CheckPostByUser(userID, postID)
	err.CheckErrWithPanic(er)

	// if the post do not belongs to the user
	if has == false {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	has, er = model.CheckCommentByPost(postID, commentID)
	err.CheckErrWithPanic(er)

	// if the comment do not belongs to the post
	if has == false {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	// delete all upvoting info of the comment
	_, er = model.CancelUpvoteCommentByComment(commentID)
	err.CheckErrWithPanic(er)

	_, er = model.CancelCommentByID(commentID)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{})
}
