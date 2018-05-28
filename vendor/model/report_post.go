package model

import (
	"database"
	"entity"
)

// NewReportPost .
func NewReportPost(userID int64, postID int64, reason string) (int64, error) {
	return database.Orm.Insert(&entity.ReportPost{
		UserID: userID,
		PostID: postID,
		Reason: reason,
	})
}
