package entity

import "time"

// Post .
type Post struct {
	ID          int64
	UserID      int64     `xorm:"notnull"`
	CategoryID  int64     `xorm:"notnull"`
	NameLibID   int64     `xorm:"notnull"`
	Content     string    `xorm:"TEXT"`
	PublishDate time.Time `xorm:"created DATETIME"`
	Like        int       `xorm:"default 0 notnull"`
}

// Posts .
type Posts []*Post
