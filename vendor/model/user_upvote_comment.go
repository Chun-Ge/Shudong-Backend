package model

import (
	"database"
	"entity"
)

// UpvoteCommentByUser ..
func UpvoteCommentByUser(userid, commentid int64) (int64, error) {
	return database.Orm.Table("UserUpvoteComment").Insert(
		&entity.UserUpvoteComment{
			UserID:    userid,
			CommentID: commentid,
		})
}

// CancelUpvoteCommentByUser ..
func CancelUpvoteCommentByUser(userid, commentid int64) (int64, error) {
	return database.Orm.Table("UserUpvoteComment").Delete(
		&entity.UserUpvoteComment{
			UserID:    userid,
			CommentID: commentid,
		})
}

// CheckCommentIfUpvoted ..
func CheckCommentIfUpvoted(userid, commentid int64) (bool, error) {
	return database.Orm.Table("UserUpvoteComment").Get(
		&entity.UserUpvoteComment{
			UserID:    userid,
			CommentID: commentid,
		})
}

// CountCommentUpvotes ..
func CountCommentUpvotes(commentid int64) (int64, error) {
	return database.Orm.Table("UserUpvoteComment").Count(
		&entity.UserUpvoteComment{
			CommentID: commentid,
		})
}
