package test

import "github.com/khu-dev/khumu-comment/model"

var CommentsData []*model.Comment

func init(){

	var id uint = 1
	myAnonymousComment := &model.Comment{
		Kind:           "anonymous",
		//AuthorUsername: "jinsu",
		Author: &model.KhumuUserSimple{Username: "jinsu"},
		ArticleID:      1,
		Content:        "테스트로 작성한 jinsu의 익명 코멘트",
		ParentID:       nil,
	}
	id++
	myNamedComment := &model.Comment{
		Kind:           "named",
		Author: &model.KhumuUserSimple{Username: "jinsu"},
		ArticleID:      1,
		Content:        "테스트로 작성한 jinsu의 기명 코멘트",
		ParentID:       nil,
	}
	id++
	othersAnonymousComment := &model.Comment{
		Kind:           "anonymous",
		Author: &model.KhumuUserSimple{Username: "somebody"},
		ArticleID:      1,
		Content:        "테스트로 작성한 somebody의 익명 코멘트",
		ParentID:       nil,
	}
	CommentsData = append(CommentsData, myAnonymousComment)
	CommentsData = append(CommentsData, myNamedComment)
	CommentsData = append(CommentsData, othersAnonymousComment)
}
