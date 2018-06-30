package service

import (
	"entity"
	"err"
	"model"
	"time"
)

// PostResponse is the struct of posts response returned to the client.
type PostResponse struct {
	PostID       int64     `json:"postId"`
	Author       string    `json:"author"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	CategoryName string    `json:"categoryName"`
	PublishDate  time.Time `json:"publishDate"`
	LikeCount    int64     `json:"likeCount"`
	CommentCount int64     `json:"commentCount"`
}

// CommentResponse is the struct of comments response returned to the client.
type CommentResponse struct {
	CommentID     int64     `json:"commentId"`
	Author        string    `json:"author"`
	RelatedPostID int64     `json:"relatedPostId"`
	Content       string    `json:"content"`
	PublishDate   time.Time `json:"publishDate"`
	LikeCount     int64     `json:"likeCount"`
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

func genSingleCommentResponse(comment *entity.Comment) *CommentResponse {
	upvoteCount, er := model.CountCommentUpvotes(comment.ID)
	err.CheckErrWithPanic(er)

	author, er := model.GetNameFromNameLibByID(comment.NameLibID)
	err.CheckErrWithPanic(er)

	return &CommentResponse{
		CommentID:     comment.ID,
		Author:        author.Name,
		RelatedPostID: comment.PostID,
		Content:       comment.Content,
		PublishDate:   comment.PublishDate,
		LikeCount:     upvoteCount,
	}
}

func genMultiCommentsResponse(comments entity.Comments) []*CommentResponse {
	ret := make([]*CommentResponse, len(comments))
	for idx, comment := range comments {
		ret[idx] = genSingleCommentResponse(comment)
	}
	return ret
}
