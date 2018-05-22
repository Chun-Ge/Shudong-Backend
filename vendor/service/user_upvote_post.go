package service

import (
	"err"
	"model"
	"response"

	"errors"

	"github.com/kataras/iris"
)

// UpvotePost ..
func UpvotePost(ctx iris.Context) {
	var affected int64 // = 0

	userid, er := ctx.Values().GetInt64("userid")
	postid, er := ctx.Params().GetInt64("postid")

	upvoted, er := model.CheckPostIfUpvoted(userid, postid)
	err.CheckErrWithPanic(er)

	if upvoted {
		affected, er = upvotePost(userid, postid)
	} else {
		affected, er = upvotePostCancel(userid, postid)
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
