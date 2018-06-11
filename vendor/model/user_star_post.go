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

// GetPostsStarredByUser ...
func GetPostsStarredByUser(userID int64) ([]*entity.UserStarPost, error) {
	retPosts := make([]*entity.UserStarPost, 0)
	err := database.Orm.Where("user_id = ?", userID).Desc("post_id").Find(&retPosts)
	return retPosts, err
}
