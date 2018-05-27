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
	var affected int64

	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postid")

	info := StarPostInfo{UserID: userID, PostID: postID}

	// TODO(alexandrali3): Check existance of user and post.

	starred, er := model.CheckPostIfStarred(info.UserID, info.PostID)
	err.CheckErrWithPanic(er)

	if !starred {
		affected, er = starPost(info.UserID, info.PostID)
	} else {
		affected, er = starPostCancel(info.UserID, info.PostID)
	}
	err.CheckErrWithPanic(er)

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
