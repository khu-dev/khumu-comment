package mapper

import (
	"github.com/khu-dev/khumu-comment/data"
	"github.com/khu-dev/khumu-comment/ent"
	"time"
)

// 기존의 Comment를 복제해서 CommentOutput을 만든다
// 기존의 Comment나 Author, Children 등 참조 X. 복제.
func CommentModelToOutput(src *ent.Comment, dest *data.CommentOutput) *data.CommentOutput {
	if src == nil {
		return nil
	}

	if dest == nil {
		dest = &data.CommentOutput{}
	}

	dest.ID = src.ID
	dest.Author = KhumuUserModelToSimpleOutput(src.Edges.Author, nil)
	if src.Edges.Article != nil {
		dest.Article = &src.Edges.Article.ID
	} else if src.Edges.StudyArticle != nil {
		dest.StudyArticle = &src.Edges.StudyArticle.ID
	}
	dest.Content = src.Content
	dest.Kind = src.Kind
	dest.State = src.State
	if src.Edges.Author != nil {
		if src.Edges.Article != nil && src.Edges.Article.Edges.Author != nil {
			dest.IsWrittenByArticleAuthor = src.Edges.Author.ID == src.Edges.Article.Edges.Author.ID
		}
	}
	// children은 그냥 빈 배열로 저장.
	// 필요한 경우
	dest.Children = []*data.CommentOutput{}

	return dest
}

func CopyCommentOutput(src *data.CommentOutput) *data.CommentOutput {
	dest := *src
	copiedAuthor := *src.Author
	dest.Author = &copiedAuthor
	dest.Children = make([]*data.CommentOutput, len(dest.Children))

	for i, child := range src.Children {
		dest.Children[i] = CopyCommentOutput(child)
	}

	return &dest
}

// Comment.CreatedAt을 바탕으로 Comment.CreatedAtExpression에 올바른 값을 입력시킨다.
func NewCreatedAtExpression(createdAt time.Time) string {
	// UTC 시간을 단순 한국시간으로 변경
	createdAtExp := "생성 시간"
	now := time.Now() // 한국 시간
	nowYear, nowMonth, nowDate := now.Date()
	createdYear, createdMonth, createdDate := createdAt.Date()
	if now.Sub(createdAt).Minutes() < 5 {
		createdAtExp = "지금"
	} else if nowYear == createdYear && nowMonth == createdMonth && nowDate == createdDate {
		createdAtExp = createdAt.Format("15:04")
	} else if nowYear == createdYear {
		createdAtExp = createdAt.Format("01/02 15:04")
	} else {
		createdAtExp = createdAt.Format("2006/01/02 15:04")
	}

	return createdAtExp
}
