package model

import (
	"database"
	"entity"
)

// UpvotePostByUser ..
func UpvotePostByUser(userid, postid int64) (int64, error) {
	return database.Orm.Table("user_upvot_post").Insert(
		&entity.UserUpvotePost{
			UserID: userid,
			PostID: postid,
		})
}

// CancelUpvotePostByUser ..
func CancelUpvotePostByUser(userid, postid int64) (int64, error) {
	return database.Orm.Table("user_upvot_post").Delete(
		&entity.UserUpvotePost{
			UserID: userid,
			PostID: postid,
		})
}

// CalcelUpvotingPostByPost
// delete all the upvoting info of the post
func CancelUpvotePostByPost(postid int64) (int64, error) {
	return database.Orm.Table("user_upvoting_post").Delete(
		&entity.UserUpvotePost{
			PostID: postid,
		})
}

// CheckPostIfUpvoted ..
func CheckPostIfUpvoted(userid, postid int64) (bool, error) {
	return database.Orm.Table("user_upvot_post").Exist(
		&entity.UserUpvotePost{
			UserID: userid,
			PostID: postid,
		})
}

// CountPostUpvotes ..
func CountPostUpvotes(postid int64) (int64, error) {
	return database.Orm.Table("user_upvot_post").Count(
		&entity.UserUpvotePost{
			PostID: postid,
		})
}
