package service

import (
	"err"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// CommentInfo ...
type CommentInfo struct {
	UserID  int64
	PostID  int64
	Comment struct {
		Content string `json:"content"`
	} `json:"comment"`
}

// CreateComment creates a new comment upon a post.
func CreateComment(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)

	info := CommentInfo{UserID: userID, PostID: postID}
	ctx.ReadJSON(&info)

	comment, er := model.NewCommentWithRandomName(info.UserID, info.PostID, info.Comment.Content)
	err.CheckErrWithPanic(er)

	author, er := model.GetNameFromNameLibByID(comment.NameLibID)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{
		"comment": iris.Map{
			"commentId":     comment.ID,
			"author":        author,
			"relatedPostId": comment.PostID,
			"content":       comment.Content,
			"likeCount":     0,
		},
	})
}

// GetCommentsOfAPost gets comments pertaining to a specific post.
func GetCommentsOfAPost(ctx iris.Context) {
	postID, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)

	comments, er := model.GetCommentsByPostID(postID)
	err.CheckErrWithPanic(er)

	ret := genMultiCommentsResponse(comments)

	response.OK(ctx, iris.Map{
		"comments": ret,
	})
}

// DeleteComment ...
// route: /post/{postid}/comments/{commentid}
// pre: the comment belongs to the post
// post: the comment has been deleted, meanwhile,
//       the upvoting info of the comment will have been deleted
// response result: 1. OK
//                  2. Forbidden: post do not belongs to the user
//                     or comment do not belongs to the post
//                  3. Unauthorized: the user is not valid
func DeleteComment(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postId")
	err.CheckErrWithPanic(er)
	commentID, er := ctx.Params().GetInt64("commentId")
	err.CheckErrWithPanic(er)

	has, er := model.CheckPostByUser(userID, postID)
	err.CheckErrWithPanic(er)

	// if the post do not belongs to the user
	if !has {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	has, er = model.CheckCommentByPost(postID, commentID)
	err.CheckErrWithPanic(er)

	// if the comment do not belongs to the post
	if !has {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	// delete all upvoting info of the comment
	_, er = model.CancelUpvoteCommentByComment(commentID)
	err.CheckErrWithPanic(er)

	_, er = model.CancelCommentByID(commentID)
	err.CheckErrWithPanic(er)

	response.OK(ctx, iris.Map{})
}
