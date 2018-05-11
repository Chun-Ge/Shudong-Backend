package entity

import "time"

// User represents users of Shudong.
type User struct {
	ID         int64
	Username   string    `xorm:"varchar(50) notnull unique"`
	Password   string    `xorm:"varchar(32) notnull"`
	CreateTime time.Time `xorm:"created DATETIME"`
	UpdateTime time.Time `xorm:"updated DATETIME"`
}
