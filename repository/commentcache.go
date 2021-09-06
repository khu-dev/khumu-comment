package repository

import (
	"github.com/khu-dev/khumu-comment/data"
)

type CommentCacheRepository interface {
	FindAllParentCommentsByArticleID(articleID int) (coms data.CommentEntities, err error)
}

type LikeCommentCacheRepository interface {
	FindAllByCommentID(commentID int) (likes data.LikeCommentEntities, err error)
}
