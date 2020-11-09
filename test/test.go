package test

import "github.com/khu-dev/khumu-comment/model"

var CommentsData map[string]*model.Comment
var UsersData map[string]*model.KhumuUserSimple

func init(){
	CommentsData = make(map[string]*model.Comment)
	UsersData = make(map[string]*model.KhumuUserSimple)

	var id int = 1
	CommentsData["JinsuAnonymousComment"] = &model.Comment{
		Kind:           "anonymous",
		//AuthorUsername: "jinsu",
		Author: &model.KhumuUserSimple{Username: "jinsu"},
		ArticleID:      1,
		Content:        "테스트로 작성한 jinsu의 익명 코멘트",
		ParentID:       nil,
	}
	id++
	CommentsData["JinsuNamedComment"] = &model.Comment{
		Kind:           "named",
		Author: &model.KhumuUserSimple{Username: "jinsu"},
		ArticleID:      1,
		Content:        "테스트로 작성한 jinsu의 기명 코멘트",
		ParentID:       nil,
	}
	id++
	CommentsData["SomebodyAnonymousComment"] = &model.Comment{
		Kind:           "anonymous",
		Author: &model.KhumuUserSimple{Username: "somebody"},
		ArticleID:      1,
		Content:        "테스트로 작성한 somebody의 익명 코멘트",
		ParentID:       nil,
	}
	UsersData["Jinsu"] = &model.KhumuUserSimple{Username: "jinsu"}
	UsersData["Somebody"] = &model.KhumuUserSimple{Username: "somebody"}
	UsersData["Puppy"] = &model.KhumuUserSimple{Username: "puppy"}
}
