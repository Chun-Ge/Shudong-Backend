package service

import (
	"model"

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
		return
	}
	upvoted, _ := model.CheckPostIfUpvoted(userid, postid)
	if upvoted {
		// upvotePost(userid, postid)
	} else {
		// upvotePostCancel(userid, postid)
	}
	if affected <= 0 {
		// Internal Server Error
	} else {
		// OK
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
