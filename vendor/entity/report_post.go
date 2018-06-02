package entity

import "time"

// ReportPost ...
type ReportPost struct {
	ID         int64
	UserID     int64     `xorm:"notnull"`
	PostID     int64     `xorm:"notnull"`
	Reason     string    `xorm:"TEXT"`
	CreateDate time.Time `xorm:"created DATETIME"`
}

// ReportPosts ...
type ReportPosts []*ReportPost
