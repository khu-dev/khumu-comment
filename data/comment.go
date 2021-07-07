package data

type CommentInput struct {
	Author       string  `json:"author"`
	Article      *int    `json:"article"`
	StudyArticle *int    `json:"study_article"`
	Parent       *int    `json:"parent"`
	Content      string  `json:"content"`
	Kind         *string `json:"kind"`
}

type CommentOutput struct {
	ID                int                    `json:"id"`
	Kind              string                 `json:"kind"`
	State             string                 `json:"state"`
	Author            *SimpleKhumuUserOutput `json:"author"`
	Article           *int                   `json:"article"`
	StudyArticle      *int                   `json:"study_article"`
	Parent            *int                   `json:"parent"`
	Content           string                 `json:"content"`
	Children          []*CommentOutput       `json:"children"`
	IsAuthor          bool                   `json:"is_author"`
	IsAuthorOfArticle bool                   `json:"is_author_of_article"`
	LikeCommentCount  int                    `json:"like_comment_count"`
	Liked             bool                   `json:"liked"`
	CreatedAt         string                 `json:"created_at"`
}

type LikeCommentInput struct {
	User    string `json:"username"`
	Comment int    `json:"comment"`
}
