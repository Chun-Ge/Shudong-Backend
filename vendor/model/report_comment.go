package model

import (
	"database"
	"entity"
)

// NewReportComment ...
func NewReportComment(userID int64, commentID int64, reason string) (int64, error) {
	return database.Orm.Insert(&entity.ReportComment{
		UserID:    userID,
		CommentID: commentID,
		Reason:    reason,
	})
}
