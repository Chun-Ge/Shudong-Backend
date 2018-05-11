package entity

import "time"

// Comment represents comments under each post. Currently, we consider only 1 level.
type Comment struct {
	ID          int64
	UserID      int64     `xorm:"notnull"`
	PostID      int64     `xorm:"notnull"`
	NameLibID   int64     `xorm:"notnull"`
	Content     string    `xorm:"TEXT"`
	PublishDate time.Time `xorm:"created DATETIME"`
	Like        int       `xorm:"default 0 notnull"`
}
