package service

import (
	"err"
	"github.com/kataras/iris"
	"middlewares"
	"model"
	"response"
)

type CommentInfo struct {
	UserID  int64
	PostID  int64
	Comment string `form:"comment"`
}

// new a comment for a post
func CreateComment(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postid")
	err.CheckErrWithPanic(er)

	info := CommentInfo{UserID: userID, PostID: postID}
	ctx.ReadForm(&info)

	comment, er := model.NewCommentWithRandomName(info.UserID, info.PostID, info.Comment)

	if er != nil {
		response.InternalServerError(ctx, iris.Map{})
		return
	}

	author, er := model.GetNameFromNameLibByID(comment.NameLibID)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{
		"comment": iris.Map{
			"commentId":     comment.ID,
			"author":        author,
			"relatedPostId": comment.PostID,
			"content":       comment.Content,
			"like_count":    0,
		},
	})
}
