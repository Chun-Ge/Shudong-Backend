package model

import (
	"database"
	"entity"
	e "err"
)

// NewPostWithRandomName create a new post with random name from namelib.
func NewPostWithRandomName(userID int64, category int64, title string, content string) (post *entity.Post, err error) {
	post = &entity.Post{
		UserID:     userID,
		CategoryID: category,
		Title:      title,
		Content:    content,
	}

	// GetRandomNameLib is in the same package model/namelib.go
	name, err := GetRandomNameLib()
	e.CheckErr(err)

	post.NameLibID = name.ID
	_, err = database.Orm.Insert(&post)

	e.CheckErr(err)
	return
}
