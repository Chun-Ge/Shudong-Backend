package model

import (
	"database"
	"entity"
	e "err"
)

// create a comment
func NewCommentWithRandomName(userID, postID int64, content string) (comment *entity.Comment, er error) {
	comment = &entity.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: content,
	}

	name, er := GetRandomNameLib()
	e.CheckErr(er)

	comment.NameLibID = name.ID
	_, er = database.Orm.Insert(&comment)
	e.CheckErr(er)

	return
}

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
