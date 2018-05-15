package service

import (
	"model"

	"github.com/kataras/iris"
)

// UpvotePost ..
func UpvotePost(ctx iris.Context) {
	// read cookie from ctx, (check login info?), then model.CheckPostIfUpvoted()
}

func upvotePost(userid, postid int64) (affected int64) {
	return model.UpvotePostByUser(userid, postid)
}

func upvotePostCancel(userid, postid int64) (affected int64) {
	return model.CancelUpvotePostByUser(userid, postid)
}
