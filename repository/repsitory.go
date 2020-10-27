package repository

import "github.com/khu-dev/khumu-comment/model"

type CommentRepository interface {
	List(opt *CommentQueryOption) []*model.Comment
	Get(id string) *model.Comment
}

type CommentQueryOption struct{
	ArticleID uint
	AuthorID string
}