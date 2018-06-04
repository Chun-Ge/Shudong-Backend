package model

import (
	"database"
	"entity"
)

// StarPostByUser ...
func StarPostByUser(userID, postID int64) (int64, error) {
	return database.Orm.Insert(
		&entity.UserStarPost{
			UserID: userID,
			PostID: postID,
		})
}

// CancelStarPostByUser ...
func CancelStarPostByUser(userID, postID int64) (int64, error) {
	return database.Orm.Delete(
		&entity.UserStarPost{
			UserID: userID,
			PostID: postID,
		})
}

// CheckPostIfStarred ...
func CheckPostIfStarred(userID, postID int64) (bool, error) {
	return database.Orm.Get(
		&entity.UserStarPost{
			UserID: userID,
			PostID: postID,
		})
}
