package service

import (
	"err"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// PostInfo ...
type PostInfo struct {
	UserID     int64
	CategoryID int64  `form:"category"`
	Title      string `form:"title"`
	Content    string `form:"content"`
}

// CreatePost creates a new post.
func CreatePost(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)

	info := PostInfo{UserID: userID}
	ctx.ReadForm(&info)
	post, er := model.NewPostWithRandomName(info.UserID, info.CategoryID, info.Title, info.Content)

	if er != nil {
		response.InternalServerError(ctx, iris.Map{})
		return
	}

	upvoteCount, er := model.CountPostUpvotes(post.ID)
	err.CheckErrWithPanic(er)

	author, er := model.GetNameFromNameLibByID(post.NameLibID)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{
		"post": iris.Map{
			"postId":       post.ID,
			"author":       author,
			"title":        post.Title,
			"content":      post.Content,
			"likeCount":    upvoteCount,
			"commentCount": 0,
		},
	})
}
