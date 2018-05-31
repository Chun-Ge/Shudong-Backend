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
	Reason    string `json:"reason"`
}

// CreateReportComment create a new report for comment.
func CreateReportComment(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	commentID, er := ctx.Params().GetInt64("commentid")
	err.CheckErrWithPanic(er)

	info := ReportCommentInfo{UserID: userID, CommentID: commentID}
	er = ctx.ReadJSON(&info)
	err.CheckErrWithPanic(er)

	// TODO(alexandrali3): Check the existence of userID and commentID.

	affected, er := model.NewReportComment(info.UserID, info.CommentID, info.Reason)
	err.CheckErrWithPanic(er)

	// Check if the report is successfully recorded.
	if affected != 1 {
		panic(errors.New("SQL Update Error"))
	} else {
		response.OK(ctx, iris.Map{})
	}
}
