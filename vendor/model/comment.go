package model

import (
	"database"
	"entity"
	e "err"
)

// create a comment
func NewCommentWithRandomName(userID, postID int64, content string) (comment *entity.Comment, er error) {
	comment = &entity.Comment{
		UserID:  userID,
		PostID:  postID,
		Content: content,
	}

	name, er := GetRandomNameLib()
	e.CheckErr(er)

	comment.NameLibID = name.ID
	_, er = database.Orm.Insert(&comment)
	e.CheckErr(er)
	return
}
