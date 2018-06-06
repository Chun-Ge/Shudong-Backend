package model

import (
	"database"
	"entity"
	e "err"
)

// NewCommentWithRandomName creates a comment.
func NewCommentWithRandomName(userID, postID int64, content string) (*entity.Comment, error) {
	comment := &entity.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: content,
	}

	name, er := GetRandomNameLib()
	e.CheckErr(er)

	comment.NameLibID = name.ID
	_, er = database.Orm.Insert(comment)
	e.CheckErr(er)

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
