package service

import (
	"err"
	"errors"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// ReportPostInfo .
type ReportPostInfo struct {
	UserID int64
	PostID int64  `form:"post"`
	Reason string `form:"reason"`
}

// CreateReportPost create a new report for post.
func CreateReportPost(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, e := ctx.Params().GetInt64("postid")
	err.CheckErrWithPanic(e)

	info := ReportPostInfo{UserID: userID, PostID: postID}
	ctx.ReadForm(&info)

	affected, e := model.NewReportPost(info.UserID, info.PostID, info.Reason)
	err.CheckErrWithPanic(e)

	// Check if report is successfully recorded.
	if affected != 1 {
		panic(errors.New("SQL Update Error"))
	} else {
		response.OK(ctx, iris.Map{})
	}
}
