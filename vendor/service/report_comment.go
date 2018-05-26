package service

import (
	"err"
	"errors"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// ReportCommentInfo .
type ReportCommentInfo struct {
	UserID    int64
	CommentID int64
	Reason    string `form:"reason"`
}

// CreateReportComment create a new report for comment.
func CreateReportComment(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	commentID, e := ctx.Params().GetInt64("commentid")
	err.CheckErrWithPanic(e)

	info := ReportCommentInfo{UserID: userID, CommentID: commentID}
	ctx.ReadForm(&info)

	affected, e := model.NewReportComment(info.UserID, info.CommentID, info.Reason)
	err.CheckErrWithPanic(e)

	// Check if report is successfully recorded.
	if affected != 1 {
		panic(errors.New("SQL Update Error"))
	} else {
		response.OK(ctx, iris.Map{})
	}
}
