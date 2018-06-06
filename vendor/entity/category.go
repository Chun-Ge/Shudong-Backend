package entity

// Category represents categories of posts.
type Category struct {
	ID   int64
	Name string `xorm:"notnull unique"`
}

// Categories ...
type Categories []*Category
