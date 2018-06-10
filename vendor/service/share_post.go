package service

import (
	"err"
	"middlewares"
	"model"
	"response"

	"errors"

	"github.com/kataras/iris"
)

// SharePost ...
func SharePost(ctx iris.Context) {
	var affected int64 // = 0

	userid := middlewares.GetUserID(ctx)
	postid, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)

	postExists, er := model.CheckPostIfExists(postid)
	err.CheckErrWithPanic(er)
	if !postExists {
		response.NotFound(ctx, iris.Map{})
		return
	}

	shared, er := model.CheckPostIfShared(userid, postid)
	err.CheckErrWithPanic(er)

	if !shared {
		affected, er = sharePost(userid, postid)
	}
	err.CheckErrWithPanic(er)

	if !shared && affected <= 0 {
		panic(errors.New("SQL Update Error"))
	} else {
		sharedCount, er := model.CountPostShared(postid)
		err.CheckErrWithPanic(er)

		response.OK(ctx, iris.Map{
			"currentUserShared":      true,
			"currentPostSharedCount": sharedCount,
		})
	}
}

func sharePost(userid, postid int64) (affected int64, er error) {
	affected, er = model.NewSharePost(userid, postid)
	return
}

// sharePostCancel ...
// this func may be never used, according to the logic of sharing a post.
func sharePostCancel(userid, postid int64) (affected int64, er error) {
	affected, er = model.CancelSharePost(userid, postid)
	return
}
