package entity

// UserUpvotePost is the mid-table for the relationship of upvoting post.
type UserUpvotePost struct {
	ID     int64
	UserID int64 `xorm:"notnull"`
	PostID int64 `xorm:"notnull"`
}
