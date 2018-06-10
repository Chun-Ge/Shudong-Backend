package model

import (
	"database"
	"entity"
)

// NewSharePost ...
func NewSharePost(userID, postID int64) (int64, error) {
	return database.Orm.Table("share_post").Insert(&entity.SharePost{
		UserID: userID,
		PostID: postID,
	})
}

// CancelSharePost ...
func CancelSharePost(userID, postID int64) (int64, error) {
	return database.Orm.Table("share_post").Delete(&entity.SharePost{
		UserID: userID,
		PostID: postID,
	})
}

// CountPostShared ...
func CountPostShared(postID int64) (int64, error) {
	return database.Orm.Table("share_post").Count(
		&entity.SharePost{
			PostID: postID,
		})
}

// CheckPostIfShared ...
func CheckPostIfShared(userID, postID int64) (bool, error) {
	return database.Orm.Table("share_post").Exist(&entity.SharePost{
		UserID: userID,
		PostID: postID,
	})
}
