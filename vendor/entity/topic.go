package entity

// Topic represents topics of names being shown.
type Topic struct {
	ID          int64
	Name        string `xorm:"notnull unique"`
	Description string `xorm:"TEXT"`
}
