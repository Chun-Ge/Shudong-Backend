package service

import (
	"err"
	"model"
	"response"

	"github.com/kataras/iris"
)

// UpvoteComment ..
func UpvoteComment(ctx iris.Context) {
	// read cookie from ctx, (check login info?), then model.CheckPostIfUpvoted()

	var (
		userid                      int64 = 1
		commentid                   int64 = 1
		affected                    int64 // = 0
		isLoggedIn                  = true
		er                          error
		callbackInternalServerError = response.GenCallbackInternalServerError(ctx)
	)

	if !isLoggedIn {
		response.Unauthorized(ctx, iris.Map{})
		return
	}
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
