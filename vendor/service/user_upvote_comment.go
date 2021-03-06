package service

import (
	"err"
	"errors"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// UpvoteComment ...
func UpvoteComment(ctx iris.Context) {
	var affected int64 // = 0

	userid := middlewares.GetUserID(ctx)
	commentid, er := ctx.Params().GetInt64("commentId")

	upvoted, er := model.CheckCommentIfUpvoted(userid, commentid)
	err.CheckErrWithPanic(er)

	if upvoted {
		affected, er = upvoteCommentCancel(userid, commentid)
	} else {
		affected, er = upvoteComment(userid, commentid)
	}
	err.CheckErrWithPanic(er)

	if affected <= 0 {
		panic(errors.New(err.SQLUpdateError))
	} else {
		upvoteCount, er := model.CountCommentUpvotes(commentid)
		err.CheckErrWithPanic(er)

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
