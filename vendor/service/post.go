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
	Post struct {
		CategoryName string `json:"categoryName"`
		Title        string `json:"title"`
		Content      string `json:"content"`
	} `json:"post"`
}

// CreatePost creates a new post.
func CreatePost(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)

	info := &PostInfo{}
	er := ctx.ReadJSON(info)
	err.CheckErrWithCallback(er, response.GenCallbackBadRequest(ctx))

	nilString := ""

	if info.Post.CategoryName == nilString ||
		info.Post.Title == nilString ||
		info.Post.Content == nilString {
		response.BadRequest(ctx, iris.Map{})
		ctx.StopExecution()
		return
	}

	categoryID, er := model.GetCategoryIDByName(info.Post.CategoryName)
	err.CheckErrWithPanic(er)
	if categoryID == -1 {
		response.NotFound(ctx, iris.Map{})
	}

	post, er := model.NewPostWithRandomName(userID, categoryID, info.Post.Title, info.Post.Content)
	err.CheckErrWithPanic(er)

	postResponse := genSinglePostResponse(post)

	response.Created(ctx, iris.Map{
		"post": postResponse,
	})
}

// DeletePost ...
// route: /post/{postid:int min(1)}
// pre: the post belongs to the user
// post: the post has been deleted, meanwhile,
//       all comments of the post will have been deleted,
//       and clear the info of upvoting the post
func DeletePost(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postId")
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

// RecentPostParam stores limit & offset for GetRecentPosts.
type RecentPostParam struct {
	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

// GetRecentPosts ...
func GetRecentPosts(ctx iris.Context) {
	param := &RecentPostParam{}
	ctx.ReadForm(param)

	recentPosts, er := model.GetRecentPosts(param.Limit, param.Offset)
	err.CheckErrWithPanic(er)

	ret := genMultiPostsResponse(recentPosts)

	response.OK(ctx, iris.Map{
		"posts": ret,
	})
}

// GetPostByID ...
func GetPostByID(ctx iris.Context) {
	postid, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)

	post, er := model.GetPostByID(postid)
	err.CheckErrWithPanic(er)

	if post == nil {
		response.NotFound(ctx, iris.Map{})
		return
	}

	postResponse := genSinglePostResponse(post)

	response.OK(ctx, iris.Map{
		"post": postResponse,
	})
}
