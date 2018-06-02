package model

import (
	"database"
	"entity"
	e "err"
)

// NewPostWithRandomName creates a new post with random name from namelib.
func NewPostWithRandomName(userID int64, category int64, title string, content string) (post *entity.Post, err error) {
	post = &entity.Post{
		UserID:     userID,
		CategoryID: category,
		Title:      title,
		Content:    content,
	}

	// GetRandomNameLib is in the same package model/namelib.go
	name, err := GetRandomNameLib()
	e.CheckErrWithPanic(err)

	post.NameLibID = name.ID
	_, err = database.Orm.Insert(post)
	e.CheckErrWithPanic(err)

	return
}

// CheckPostByUser checks a post by userid and postid
func CheckPostByUser(userID, postID int64) (bool, error) {
	return database.Orm.Table("post").Exist(
		&entity.Post{
			ID:     postID,
			UserID: userID,
		})
}

// CancelPostByID ...
func CancelPostByID(postID int64) (int64, error) {
	return database.Orm.Table("post").Delete(
		&entity.Post{
			ID: postID,
		})
}
