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
	PostID int64
	Reason string `form:"reason"`
}

// CreateReportPost create a new report for post.
func CreateReportPost(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postid")
	err.CheckErrWithPanic(er)

	info := ReportPostInfo{UserID: userID, PostID: postID}
	er = ctx.ReadForm(&info)
	err.CheckErrWithPanic(er)

	affected, er := model.NewReportPost(info.UserID, info.PostID, info.Reason)
	err.CheckErrWithPanic(er)

	// Check if the report is successfully recorded.
	if affected != 1 {
		panic(errors.New("SQL Update Error"))
	} else {
		response.OK(ctx, iris.Map{})
	}
}
