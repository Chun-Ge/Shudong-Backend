package entity

// UserUpvoteComment is the mid-table for the relationship of upvoting comment.
type UserUpvoteComment struct {
	ID        int64
	UserID    int64 `xorm:"notnull"`
	CommentID int64 `xorm:"notnull"`
}
