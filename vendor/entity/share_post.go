package entity

import "time"

// SharePost ...
type SharePost struct {
	ID         int64
	UserID     int64     `xorm:"notnull"`
	PostID     int64     `xorm:"notnull"`
	CreateDate time.Time `xorm:"created DATETIME"`
}

// SharePosts ...
type SharePosts []*SharePost
