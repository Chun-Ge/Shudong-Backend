package entity

import "time"

// AuthCode for reset password
// foreigh key : UserID reference to the User
type AuthCode struct {
	ID         int64
	UserID     int64     `xorm:"unique notnull"`
	Code       string    `xorm:"notnull"`
	UpdateTime time.Time `xorm:"updated DATETIME"`
}
