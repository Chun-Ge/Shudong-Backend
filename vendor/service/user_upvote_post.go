package service

import (
	"err"
	"middlewares"
	"model"
	"response"

	"errors"

	"github.com/kataras/iris"
)

// UpvotePost ...
func UpvotePost(ctx iris.Context) {
	var affected int64 // = 0

	userid := middlewares.GetUserID(ctx)
	postid, er := ctx.Params().GetInt64("postId")

	upvoted, er := model.CheckPostIfUpvoted(userid, postid)
	err.CheckErrWithPanic(er)

	if upvoted {
		affected, er = upvotePostCancel(userid, postid)
	} else {
		affected, er = upvotePost(userid, postid)
	}
	err.CheckErrWithPanic(er)

	if affected <= 0 {
		panic(errors.New("SQL Update Error"))
	} else {
		upvoteCount, er := model.CountPostUpvotes(postid)
		err.CheckErrWithPanic(er)

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
