package entity

// UserStarPost is the mid-table for the relationship of starring post.
type UserStarPost struct {
	ID     int64
	UserID int64 `xorm:"notnull"`
	PostID int64 `xorm:"notnull"`
}
