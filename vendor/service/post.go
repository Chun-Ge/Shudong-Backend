package service

import (
	"err"
	"fmt"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// PostInfo ...
type PostInfo struct {
	UserID     int64
	CategoryID int64  `json:"category"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

// CreatePost creates a new post.
func CreatePost(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)

	info := &PostInfo{UserID: userID}
	ctx.ReadJSON(info)
	fmt.Println(info.UserID, info.CategoryID, info.Title, info.Content)
	post, er := model.NewPostWithRandomName(info.UserID, info.CategoryID, info.Title, info.Content)
	err.CheckErrWithPanic(er)

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

// DeletePost .
// route: /post/{postid:int min(1)}
// pre: the post belongs to the user
// post: the post has been deleted, meanwhile,
//       all comments of the post will have been deleted,
//       and clear the info of upvoting the post
func DeletePost(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postid")
	err.CheckErrWithPanic(er)

	has, er := model.CheckPostByUser(userID, postID)
	err.CheckErrWithPanic(er)

	// if the post do not belong to the user
	if has == false {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	// delete the upvoting info of the post
	_, er = model.CancelUpvotePostByPost(postID)
	err.CheckErrWithPanic(er)

	// delete all comments of the post
	_, er = model.CancelCommentByPost(postID)
	err.CheckErrWithPanic(er)

	// delete the post
	_, er = model.CancelPostByID(postID)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{})
}
