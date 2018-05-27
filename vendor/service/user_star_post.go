package service

import (
	"err"
	"middlewares"
	"model"
	"response"

	"errors"

	"github.com/kataras/iris"
)

// StarPostInfo .
type StarPostInfo struct {
	UserID int64
	PostID int64
}

// StarPost .
func StarPost(ctx iris.Context) {
	var affected int64 // = 0

	userID := middlewares.GetUserID(ctx)
	postID, e := ctx.Params().GetInt64("postid")

	info := StarPostInfo{UserID: userID, PostID: postID}

	starred, e := model.CheckPostIfStarred(info.UserID, info.PostID)
	err.CheckErrWithPanic(e)

	if starred {
		affected, e = starPost(info.UserID, info.PostID)
	} else {
		affected, e = starPostCancel(info.UserID, info.PostID)
	}
	err.CheckErrWithPanic(e)

	starred = !starred

	if affected != 1 {
		panic(errors.New("SQL Update Error"))
	} else {
		response.OK(ctx, iris.Map{
			"currentUserStarred": starred,
		})
	}
}

func starPost(userID, postID int64) (affected int64, er error) {
	affected, er = model.StarPostByUser(userID, postID)
	return
}

func starPostCancel(userID, postID int64) (affected int64, er error) {
	affected, er = model.CancelStarPostByUser(userID, postID)
	return
}
