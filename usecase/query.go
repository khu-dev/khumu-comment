package usecase

import "github.com/khu-dev/khumu-comment/ent"

// 댓글 조회 시 기본적으로 필요한 query문을 추가한다.
// 댓글 작성자
// 게시물과 게시물의 작성자
// 스터디 게시물과 스터디 게시물의 작성자
// 대댓글도 마찬가지 과정
func appendQueryForComment(query *ent.CommentQuery) *ent.CommentQuery {
	query.
		// 댓글 작성자
		WithAuthor().
		WithArticle(func(query *ent.ArticleQuery) {
			query.WithAuthor()
		}).
		WithStudyArticle(func(query *ent.StudyArticleQuery) {
			query.WithAuthor()
		}).
		// 대댓글
		WithChildren(
			func(query *ent.CommentQuery) {
				query.
					WithAuthor().
					WithArticle(func(query *ent.ArticleQuery) {
						query.WithAuthor()
					}).
					WithStudyArticle(func(query *ent.StudyArticleQuery) {
						query.WithAuthor()
					})
			},
		)

	return query
}
