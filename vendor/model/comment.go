package model

import (
	"database"
	"entity"
)

// NewCommentWithRandomName creates a comment.
func NewCommentWithRandomName(userID, postID int64, content string) (*entity.Comment, error) {
	comment := &entity.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: content,
	}

	name, er := GetRandomNameLib()

	comment.NameLibID = name.ID
	_, er = database.Orm.Insert(comment)

	return comment, er
}

// CheckCommentByPost checks whether the comment belongs to the post
func CheckCommentByPost(postID, commentID int64) (bool, error) {
	return database.Orm.Table("comment").Exist(
		&entity.Comment{
			ID:     commentID,
			PostID: postID,
		})
}

// CheckCommentByID checks the existence of a comment by commentID.
func CheckCommentByID(commentID int64) (bool, error) {
	return database.Orm.Exist(
		&entity.Comment{
			ID: commentID,
		})
}

// CancelCommentByPost deletes all comments of the post
func CancelCommentByPost(postID int64) (int64, error) {
	return database.Orm.Table("comment").Delete(
		&entity.Comment{
			PostID: postID,
		})
}

// CancelCommentByID deletes a comment by id
func CancelCommentByID(commentID int64) (int64, error) {
	return database.Orm.Table("comment").Delete(
		&entity.Comment{
			ID: commentID,
		})
}

// CountCommentsOfPost ...
func CountCommentsOfPost(postID int64) (int64, error) {
	return database.Orm.Table("comment").Count(
		&entity.Comment{
			PostID: postID,
		})
}

// GetCommentsByPostID ...
func GetCommentsByPostID(postID int64) (entity.Comments, error) {
	comments := make(entity.Comments, 0)
	er := database.Orm.Where("post_id = ?", postID).Asc("publish_date").Find(&comments)
	return comments, er
}

// GetCommentsByUserID ...
func GetCommentsByUserID(userID int64) (entity.Comments, error) {
	comments := make(entity.Comments, 0)
	er := database.Orm.Where("user_id = ?", userID).Desc("publish_date").Find(&comments)
	return comments, er
}
