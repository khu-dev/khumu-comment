package mapper

import (
    "github.com/khu-dev/khumu-comment/config"
    "github.com/khu-dev/khumu-comment/data"
    "github.com/khu-dev/khumu-comment/ent"
    "time"
)

func CommentModelToOutput(src *ent.Comment, dest *data.CommentOutput) *data.CommentOutput {
    if src == nil {
        return nil
    }

    if dest == nil {
        dest = &data.CommentOutput{}
    }

    dest.Id = src.ID
    dest.Author = KhumuUserModelToSimpleOutput(src.Edges.Author, nil)
    dest.Article = &src.Edges.Article.ID
    dest.StudyArticle = nil
    dest.Content = src.Content
    dest.Children = []*data.CommentOutput{}


    return dest
}

// Comment.CreatedAt을 바탕으로 Comment.CreatedAtExpression에 올바른 값을 입력시킨다.
func NewCreatedAtExpression(createdAt time.Time) string {
	// UTC 시간을 단순 한국시간으로 변경
    createdAtExp := "생성 시간"
	now := time.Now().In(config.Location) // now는 근데 기본적으로 UTC긴한듯.
	nowYear, nowMonth, nowDate := now.Date()
	//log.Println(c.CreatedAt.In(repository.Location).Format("2006/01/02 15:04")) // => 한국시간대로 잘 나옴.
	localCreatedAt := createdAt.In(config.Location)
	createdYear, createdMonth, createdDate := createdAt.Date()
	if now.Sub(createdAt).Minutes() < 5 {
		createdAtExp = "지금"
	} else if nowYear == createdYear && nowMonth == createdMonth && nowDate == createdDate {
		createdAtExp = localCreatedAt.Format("15:04")
	} else if nowYear == createdYear {
		createdAtExp = localCreatedAt.Format("01/02 15:04")
	} else {
		createdAtExp = localCreatedAt.Format("2006/01/02 15:04")
	}

	return createdAtExp
}