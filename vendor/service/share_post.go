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

	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)

	postExists, er := model.CheckPostIfExists(postID)
	err.CheckErrWithPanic(er)
	if !postExists {
		response.NotFound(ctx, iris.Map{})
		return
	}

	shared, er := model.CheckPostIfShared(userID, postID)
	err.CheckErrWithPanic(er)

	if !shared {
		affected, er = sharePost(userID, postID)
	}
	err.CheckErrWithPanic(er)

	if !shared && affected <= 0 {
		panic(errors.New("SQL Update Error"))
	} else {
		sharedCount, er := model.CountPostShared(postID)
		err.CheckErrWithPanic(er)

		response.OK(ctx, iris.Map{
			"currentUserShared":      true,
			"currentPostSharedCount": sharedCount,
		})
	}
}

func sharePost(userID, postID int64) (affected int64, er error) {
	affected, er = model.NewSharePost(userID, postID)
	return
}

// sharePostCancel ...
// this func may be never used, according to the logic of sharing a post.
func sharePostCancel(userID, postID int64) (affected int64, er error) {
	affected, er = model.CancelSharePost(userID, postID)
	return
}
