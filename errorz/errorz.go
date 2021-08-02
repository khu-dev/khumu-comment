package errorz

import "errors"

var (
	ErrResourceNotFound = errors.New("해당 리소스를 찾을 수 없습니다")
	ErrUnauthorized     = errors.New("요청을 수행할 권한이 없습니다")
	ErrSelfLikeComment  = errors.New("본인의 댓글은 좋아요할 수 없습니다")
	ErrNoArticleIDInput = errors.New("게시물 ID를 입력하십시오")
)
