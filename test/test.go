package test

import "github.com/khu-dev/khumu-comment/model"

var CommentsData map[string]*model.Comment
var UsersData map[string]*model.KhumuUserSimple

func init(){
	CommentsData = make(map[string]*model.Comment)
	UsersData = make(map[string]*model.KhumuUserSimple)

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
	CommentsData["AnonymousJinsuComment"] = myAnonymousComment
	CommentsData["NamedJinsuComment"] = myNamedComment
	CommentsData["AnonymousSomebodyComment"] = othersAnonymousComment

	userJinsu := &model.KhumuUserSimple{Username: "jinsu"}
	userSomebody := &model.KhumuUserSimple{Username: "somebody"}
	userPuppy := &model.KhumuUserSimple{Username: "puppy"}
	UsersData["jinsu"] = userJinsu
	UsersData["somebody"] = userSomebody
	UsersData["puppy"] = userPuppy
}
