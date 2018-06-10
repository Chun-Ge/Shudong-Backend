package service

import (
	"err"
	"errors"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// ReportPostInfo ...
type ReportPostInfo struct {
	UserID int64
	PostID int64
	Reason string `json:"reason"`
}

// CreateReportPost creates a new report for post.
func CreateReportPost(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)

	info := ReportPostInfo{UserID: userID, PostID: postID}
	er = ctx.ReadJSON(&info)
	err.CheckErrWithCallback(er, response.GenCallbackBadRequest(ctx))

	// TODO(alexandrali3): Check the existence of userID and postID.

	affected, er := model.NewReportPost(info.UserID, info.PostID, info.Reason)
	err.CheckErrWithPanic(er)

	// Check if the report is successfully recorded.
	if affected != 1 {
		panic(errors.New("SQL Update Error"))
	} else {
		response.OK(ctx, iris.Map{})
	}
}
