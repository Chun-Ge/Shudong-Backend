package model

import (
	"database"
	"entity"
)

// check whether the comment belongs to the post
func CheckCommentByPost(postID, commentID int64) (bool, error) {
	return database.Orm.Table("comment").Exist(
		&entity.Comment{
			ID:     commentID,
			PostID: postID,
		})
}

// delete all comments of the post
func CancelCommentByPost(postID int64) (int64, error) {
	return database.Orm.Table("comment").Delete(
		&entity.Comment{
			PostID: postID,
		})
}

// delete comment by id
func CancelCommentByID(commentID int64) (int64, error) {
	return database.Orm.Table("comment").Delete(
		&entity.Comment{
			ID: commentID,
		})
}
