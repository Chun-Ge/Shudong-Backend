package service

import (
	"err"
	"model"
	"response"

	"github.com/kataras/iris"
)

// UpvotePost ..
func UpvotePost(ctx iris.Context) {
	var (
		affected                    int64 // = 0
		callbackInternalServerError = response.GenCallbackInternalServerError(ctx)
	)

	userid, er := ctx.Values().GetInt64("userid")
	postid, er := ctx.Params().GetInt64("postid")

	upvoted, er := model.CheckPostIfUpvoted(userid, postid)
	err.CheckErrWithCallback(er, callbackInternalServerError)

	if upvoted {
		affected, er = upvotePost(userid, postid)
	} else {
		affected, er = upvotePostCancel(userid, postid)
	}
	err.CheckErrWithCallback(er, callbackInternalServerError)

	if affected <= 0 {
		callbackInternalServerError()
	} else {
		upvoteCount, er := model.CountPostUpvotes(postid)
		err.CheckErrWithCallback(er, callbackInternalServerError)

		response.OK(ctx, iris.Map{
			"currentUserLike":  !upvoted,
			"currentLikeCount": upvoteCount,
		})
	}
}

func upvotePost(userid, postid int64) (affected int64, er error) {
	affected, er = model.UpvotePostByUser(userid, postid)
	return
}

func upvotePostCancel(userid, postid int64) (affected int64, er error) {
	affected, er = model.CancelUpvotePostByUser(userid, postid)
	return
}
