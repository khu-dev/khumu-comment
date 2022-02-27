package cache

import (
	"github.com/khu-dev/khumu-comment/data"
)

type CommentCacheRepository interface {
	FindAllParentCommentsByArticleID(articleID int) (coms data.CommentEntities, err error)
	SetCommentsByArticleID(articleID int, coms data.CommentEntities)
	SetCommentCountByArticleID(articleID, cnt int)
	Count(articleID int) (cnt int, err error)
}

type LikeCommentCacheRepository interface {
	FindAllByCommentID(commentID int) (likes data.LikeCommentEntities, err error)
	SetLikesByCommentID(commentID int, likes data.LikeCommentEntities)
}
