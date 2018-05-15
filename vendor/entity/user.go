package entity

import "time"

// User represents users of Shudong.
// People use Email and Password to sign in, then the Server will return the related Userid,
// which is stored by the Frontend automatically and used for [/users/{userid}].
// Note: People don't need to remember the Userid.
type User struct {
	ID       int64
	Email    string `xorm:"notnull unique"`
	Password string `xorm:"varchar(32) notnull"`
	// Userid     string    `xorm:"varchar(50)"`
	CreateTime time.Time `xorm:"created DATETIME"`
	UpdateTime time.Time `xorm:"updated DATETIME"`
}
