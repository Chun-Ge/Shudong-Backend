package model

import (
	"database"
	"entity"
)

// check a post by userid and postid
func CheckPostByUser(userID, postID int64) (bool, error) {
	return database.Orm.Table("post").Exist(
		&entity.Post{
			ID:     postID,
			UserID: userID,
		})
}

// DeletePostByPost
func CancelPostByID(postID int64) (int64, error) {
	return database.Orm.Table("post").Delete(
		&entity.Post{
			ID: postID,
		})
}
