package errorz

import (
	goerrors "errors"
	"github.com/pkg/errors"
)

var (
	ErrResourceNotFound = goerrors.New("해당 리소스를 찾을 수 없습니다")
	ErrUnauthorized     = goerrors.New("요청을 수행할 권한이 없습니다")
	ErrBadRequest       = goerrors.New("잘못된 요청입니다")
	ErrNoRequiredField  = goerrors.New("필수 입력값을 입력해주세요")

	ErrSelfLikeComment  = errors.WithMessage(ErrBadRequest, "본인의 댓글은 좋아요할 수 없습니다")
	ErrNoArticleIDInput = errors.WithMessage(ErrNoRequiredField, "게시물 ID를 입력하십시오")
	ErrNoCommentIDInput = errors.WithMessage(ErrNoRequiredField, "댓글 ID를 입력하십시오")
	ErrNoCommentAuthor  = errors.WithMessage(ErrNoRequiredField, "댓글 작성자를 입력하십시오")
	ErrWrongArticle     = errors.WithMessage(ErrBadRequest, "잘못된 게시글입니다")
)
