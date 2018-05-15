package model

import (
	"database"
	"entity"
	"err"
)

// UpvotePostByUser ..
func UpvotePostByUser(userid, postid int64) (affected int64) {
	newUpvote := &entity.UserUpvotePost{
		UserID: userid,
		PostID: postid,
	}
	affected, er := database.Orm.Table("UserUpvotePost").Insert(newUpvote)
	err.CheckErr(er)
	return
}

// CancelUpvotePostByUser ..
func CancelUpvotePostByUser(userid, postid int64) (affected int64) {
	delUpvote := &entity.UserUpvotePost{
		UserID: userid,
		PostID: postid,
	}
	affected, er := database.Orm.Table("UserUpvotePost").Delete(delUpvote)
	err.CheckErr(er)
	return
}

// CheckPostIfUpvoted ..
func CheckPostIfUpvoted(userid, postid int64) (has bool) {
	searchEntry := &entity.UserUpvotePost{
		UserID: userid,
		PostID: postid,
	}
	has, er := database.Orm.Table("UserUpvotePost").Get(searchEntry)
	err.CheckErr(er)
	return
}
