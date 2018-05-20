package model

import (
	"database"
	"entity"
)

// UpvotePostByUser ..
func UpvotePostByUser(userid, postid int64) (int64, error) {
	return database.Orm.Table("UserUpvotePost").Insert(
		&entity.UserUpvotePost{
			UserID: userid,
			PostID: postid,
		})
}

// CancelUpvotePostByUser ..
func CancelUpvotePostByUser(userid, postid int64) (int64, error) {
	return database.Orm.Table("UserUpvotePost").Delete(
		&entity.UserUpvotePost{
			UserID: userid,
			PostID: postid,
		})
}

// CheckPostIfUpvoted ..
func CheckPostIfUpvoted(userid, postid int64) (bool, error) {
	return database.Orm.Table("UserUpvotePost").Get(
		&entity.UserUpvotePost{
			UserID: userid,
			PostID: postid,
		})
}

// CountPostUpvotes ..
func CountPostUpvotes(postid int64) (int64, error) {
	return database.Orm.Table("UserUpvotePost").Count(
		&entity.UserUpvotePost{
			PostID: postid,
		})
}
