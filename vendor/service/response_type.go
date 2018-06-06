package service

import (
	"entity"
	"err"
	"model"
	"time"
)

// PostResponse is the struct of response returned to the client.
type PostResponse struct {
	PostID       int64     `json:"postid"`
	Author       string    `json:"author"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CategoryName string    `json:"categoryName"`
	PublishDate  time.Time `json:"publishDate"`
	LikeCount    int64     `json:"likeCount"`
	CommentCount int64     `json:"commentCount"`
}

func genSinglePostResponse(post *entity.Post) *PostResponse {
	upvoteCount, er := model.CountPostUpvotes(post.ID)
	err.CheckErrWithPanic(er)

	commentCount, er := model.CountCommentsOfPost(post.ID)
	err.CheckErrWithPanic(er)

	author, er := model.GetNameFromNameLibByID(post.NameLibID)
	err.CheckErrWithPanic(er)

	categoryName, er := model.GetCategoryNameByID(post.CategoryID)
	err.CheckErrWithPanic(er)

	return &PostResponse{
		PostID:       post.ID,
		Author:       author.Name,
		Title:        post.Title,
		Content:      post.Content,
		CategoryName: categoryName,
		PublishDate:  post.PublishDate,
		LikeCount:    upvoteCount,
		CommentCount: commentCount,
	}
}

func genMultiPostsResponse(posts entity.Posts) []*PostResponse {
	ret := make([]*PostResponse, len(posts))
	for idx, post := range posts {
		ret[idx] = genSinglePostResponse(post)
	}
	return ret
}
