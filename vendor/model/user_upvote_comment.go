package model

import (
	"database"
	"entity"
)

// UpvoteCommentByUser ...
func UpvoteCommentByUser(userid, commentid int64) (int64, error) {
	return database.Orm.Table("user_upvote_comment").Insert(
		&entity.UserUpvoteComment{
			UserID:    userid,
			CommentID: commentid,
		})
}

// CancelUpvoteCommentByUser ...
func CancelUpvoteCommentByUser(userid, commentid int64) (int64, error) {
	return database.Orm.Table("user_upvote_comment").Delete(
		&entity.UserUpvoteComment{
			UserID:    userid,
			CommentID: commentid,
		})
}

// CancelUpvoteCommentByComment ...
// delete all upvoting info of the comment
func CancelUpvoteCommentByComment(commentID int64) (int64, error) {
	return database.Orm.Table("user_upvote_comment").Delete(
		&entity.UserUpvoteComment{
			CommentID: commentID,
		})
}

// CheckCommentIfUpvoted ...
func CheckCommentIfUpvoted(userid, commentid int64) (bool, error) {
	return database.Orm.Table("user_upvote_comment").Exist(
		&entity.UserUpvoteComment{
			UserID:    userid,
			CommentID: commentid,
		})
}

// CountCommentUpvotes ...
func CountCommentUpvotes(commentID int64) (int64, error) {
	return database.Orm.Table("user_upvote_comment").Count(
		&entity.UserUpvoteComment{
			CommentID: commentID,
		})
}
