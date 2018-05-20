package service

import (
	"model"
	"response"

	"github.com/kataras/iris"
)

// UpvotePost ..
func UpvotePost(ctx iris.Context) {
	// read cookie from ctx, (check login info?), then model.CheckPostIfUpvoted()

	var (
		userid     int64 = 1
		postid     int64 = 1
		affected   int64 // = 0
		isLoggedIn = true
		// er         error
	)

	if !isLoggedIn {
		// Unauthorized
		response.Unauthorized(ctx, iris.Map{})
		return
	}
	upvoted, _ := model.CheckPostIfUpvoted(userid, postid)
	if upvoted {
		upvotePost(userid, postid)
	} else {
		upvotePostCancel(userid, postid)
	}
	if affected <= 0 {
		// Internal Server Error
		response.InternalServerError(ctx, iris.Map{})
	} else {
		// OK
		upvoteCount, _ := model.CountUpvotes(postid)
		response.OK(ctx, iris.Map{
			"currentUserLike":  !upvoted,
			"currentLikeCount": upvoteCount,
		})
	}
}

func upvotePost(userid, postid int64) (affected int64) {
	affected, _ = model.UpvotePostByUser(userid, postid)
	return
}

func upvotePostCancel(userid, postid int64) (affected int64) {
	affected, _ = model.CancelUpvotePostByUser(userid, postid)
	return
}
