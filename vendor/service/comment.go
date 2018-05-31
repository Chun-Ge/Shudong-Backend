package service

import (
	"err"
	"middlewares"
	"model"
	"response"

	"github.com/kataras/iris"
)

// CommentInfo .
type CommentInfo struct {
	UserID  int64
	PostID  int64
	Comment string `json:"comment"`
}

// CreateComment creates a new comment upon a post.
func CreateComment(ctx iris.Context) {
	userID := middlewares.GetUserID(ctx)
	postID, er := ctx.Params().GetInt64("postid")
	err.CheckErrWithPanic(er)

	info := CommentInfo{UserID: userID, PostID: postID}
	ctx.ReadJSON(&info)

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
	postID, er := ctx.Params().GetInt64("postid")
	err.CheckErrWithPanic(er)
	commentID, er := ctx.Params().GetInt64("commentid")
	err.CheckErrWithPanic(er)

	has, er := model.CheckPostByUser(userID, postID)
	err.CheckErrWithPanic(er)

	// if the post do not belongs to the user
	if has == false {
		response.Forbidden(ctx, iris.Map{})
		return
	}

	has, er = model.CheckCommentByPost(postID, commentID)
	err.CheckErrWithPanic(er)

	// if the comment do not belongs to the post
	if has == false {
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
