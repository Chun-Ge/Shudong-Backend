package entity

import "time"

// ReportComment ...
type ReportComment struct {
	ID         int64
	UserID     int64     `xorm:"notnull"`
	CommentID  int64     `xorm:"notnull"`
	Reason     string    `xorm:"TEXT"`
	CreateDate time.Time `xorm:"created DATETIME"`
}

// ReportComments ...
type ReportComments []*ReportComment
