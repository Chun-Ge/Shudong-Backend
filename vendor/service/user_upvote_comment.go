package service

import (
	"err"
	"model"
	"response"

	"github.com/kataras/iris"
)

// UpvoteComment ..
func UpvoteComment(ctx iris.Context) {
	var (
		affected                    int64 // = 0
		callbackInternalServerError = response.GenCallbackInternalServerError(ctx)
	)

	userid, er := ctx.Values().GetInt64("userid")
	commentid, er := ctx.Params().GetInt64("commentid")

	upvoted, er := model.CheckCommentIfUpvoted(userid, commentid)
	err.CheckErrWithCallback(er, callbackInternalServerError)

	if upvoted {
		affected, er = upvoteComment(userid, commentid)
	} else {
		affected, er = upvoteCommentCancel(userid, commentid)
	}
	err.CheckErrWithCallback(er, callbackInternalServerError)

	if affected <= 0 {
		callbackInternalServerError()
	} else {
		upvoteCount, er := model.CountCommentUpvotes(commentid)
		err.CheckErrWithCallback(er, callbackInternalServerError)

		response.OK(ctx, iris.Map{
			"currentUserLike":  !upvoted,
			"currentLikeCount": upvoteCount,
		})
	}
}

func upvoteComment(userid, commentid int64) (affected int64, er error) {
	affected, er = model.UpvoteCommentByUser(userid, commentid)
	return
}

func upvoteCommentCancel(userid, commentid int64) (affected int64, er error) {
	affected, er = model.CancelUpvoteCommentByUser(userid, commentid)
	return
}
