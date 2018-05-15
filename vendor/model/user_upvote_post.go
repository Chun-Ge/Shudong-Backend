package model

import (
	"database"
	"entity"
)

// UpvotePostByUser ..
func UpvotePostByUser(userid, postid int64) (int64, error) {
	newUpvote := &entity.UserUpvotePost{
		UserID: userid,
		PostID: postid,
	}
	return database.Orm.Table("UserUpvotePost").Insert(newUpvote)
}

// CancelUpvotePostByUser ..
func CancelUpvotePostByUser(userid, postid int64) (int64, error) {
	delUpvote := &entity.UserUpvotePost{
		UserID: userid,
		PostID: postid,
	}
	return database.Orm.Table("UserUpvotePost").Delete(delUpvote)
}

// CheckPostIfUpvoted ..
func CheckPostIfUpvoted(userid, postid int64) (bool, error) {
	searchEntry := &entity.UserUpvotePost{
		UserID: userid,
		PostID: postid,
	}
	return database.Orm.Table("UserUpvotePost").Get(searchEntry)
}

// CountUpvotes ..
func CountUpvotes(postid int64) (int64, error) {
	return database.Orm.Table("UserUpvotePost").Count(
		&entity.UserUpvotePost{
			PostID: postid,
		})
}
