package data

import (
	"encoding/json"
	"github.com/khu-dev/khumu-comment/ent"
)

type CommentInput struct {
	Author       string  `json:"author"`
	Article      *int    `json:"article"`
	StudyArticle *int    `json:"study_article"`
	Parent       *int    `json:"parent"`
	Content      string  `json:"content"`
	Kind         *string `json:"kind"`
}

type CommentOutput struct {
	ID                       int                 `json:"id"`
	Kind                     string              `json:"kind"`
	State                    string              `json:"state"`
	Author                   *SimpleKhumuUserDto `json:"author"`
	Article                  *int                `json:"article"`
	StudyArticle             *int                `json:"study_article"`
	Parent                   *int                `json:"parent"`
	Content                  string              `json:"content"`
	Children                 []*CommentOutput    `json:"children"`
	IsAuthor                 bool                `json:"is_author"`
	IsWrittenByArticleAuthor bool                `json:"is_written_by_article_author"`
	LikeCommentCount         int                 `json:"like_comment_count"`
	Liked                    bool                `json:"liked"`
	CreatedAt                string              `json:"created_at"`
}

type LikeCommentInput struct {
	User    string `json:"username"`
	Comment int    `json:"comment"`
}

type CommentEntities []*ent.Comment

type LikeCommentEntities []*ent.LikeComment

// MarshalBinary -
func (cl *CommentEntities) MarshalBinary() ([]byte, error) {
	return json.Marshal(cl)
}

// UnmarshalBinary -
func (cl *CommentEntities) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, cl); err != nil {
		return err
	}

	return nil
}

func (cl CommentEntities) GetTotalLength() int {
	length := 0
	for _, parent := range cl {
		length += len(parent.Edges.Children)
	}
	length += len(cl)

	return length
}

func (cl CommentEntities) GetParentsLength() int {
	return len(cl)
}

// MarshalBinary -
func (lcl *LikeCommentEntities) MarshalBinary() ([]byte, error) {
	return json.Marshal(lcl)
}

// UnmarshalBinary -
func (lcl *LikeCommentEntities) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, lcl); err != nil {
		return err
	}

	return nil
}

// UnmarshalBinary -
func (lcl LikeCommentEntities) GetLiked(username string) bool {
	for _, like := range lcl {
		if like.Edges.LikedBy.ID == username {
			return true
		}
	}

	return false
}
