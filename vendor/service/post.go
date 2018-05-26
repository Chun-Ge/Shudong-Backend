package service

import (
	"err"
	"github.com/kataras/iris"
	"middlewares"
	"model"
	"response"
)

type PostInfo struct {
	UserID     int64
	CategoryID int64  `form:"category"`
	Title      string `form:"title"`
	Content    string `form:"content"`
}

// new a post
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
			"summary":      "new post summary", // bug!!
			"content":      post.Content,
			"likeCount":    upvoteCount,
			"commentCount": 0,
		},
	})
}

// delete a post
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
