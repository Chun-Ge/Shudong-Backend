package model

import (
	"database"
	"entity"
)

// NewPostWithRandomName creates a new post with random name from namelib.
func NewPostWithRandomName(userID int64, category int64, title string, content string) (*entity.Post, error) {
	post := &entity.Post{
		UserID:     userID,
		CategoryID: category,
		Title:      title,
		Content:    content,
	}

	// GetRandomNameLib is in the same package model/namelib.go
	name, er := GetRandomNameLib()

	post.NameLibID = name.ID
	_, er = database.Orm.Insert(post)

	return post, er
}

// CheckPostByUser checks a post by userid and postid
func CheckPostByUser(userID, postID int64) (bool, error) {
	return database.Orm.Table("post").Exist(
		&entity.Post{
			ID:     postID,
			UserID: userID,
		})
}

// CheckPostByID checks the existence of a post by postID.
func CheckPostByID(postID int64) (bool, error) {
	return database.Orm.Exist(
		&entity.Post{
			ID: postID,
		})
}

// CancelPostByID ...
func CancelPostByID(postID int64) (int64, error) {
	return database.Orm.Table("post").Delete(
		&entity.Post{
			ID: postID,
		})
}

// GetRecentPosts ...
func GetRecentPosts(limit, offset int, categoryID int64) (entity.Posts, error) {
	var err error
	recentPosts := make(entity.Posts, 0)
	if categoryID == -1 {
		err = database.Orm.Table("post").Desc("publish_date").Find(&recentPosts)
	} else {
		err = database.Orm.Table("post").Where(
			"category_id = ?", categoryID).Desc("publish_date").Find(&recentPosts)
	}
	endIdx := offset + limit
	lenPosts := len(recentPosts)
	if offset >= lenPosts {
		return make(entity.Posts, 0), err
	}
	if endIdx > lenPosts {
		endIdx = lenPosts
	}
	retPosts := recentPosts[offset:endIdx]
	return retPosts, err
}

// GetPostByID ...
func GetPostByID(postid int64) (*entity.Post, error) {
	ret := &entity.Post{
		ID: postid,
	}
	has, err := database.Orm.Table("post").Get(ret)
	if !has {
		return nil, err
	}
	return ret, err
}

// CheckPostIfExists ...
func CheckPostIfExists(postID int64) (bool, error) {
	return database.Orm.Table("post").Exist(&entity.Post{
		ID: postID,
	})
}
